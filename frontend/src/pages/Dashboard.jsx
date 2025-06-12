import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'

// Sample task data
const sampleTasks = [
  { id: 1, title: 'Complete project proposal', status: 'In Progress', priority: 'High', dueDate: '2024-03-20' },
  { id: 2, title: 'Review pull requests', status: 'Pending', priority: 'Medium', dueDate: '2024-03-18' },
  { id: 3, title: 'Update documentation', status: 'Completed', priority: 'Low', dueDate: '2024-03-15' },
]

const menuItems = [
  {
    label: 'Tasks',
    icon: (active) => (
      <svg className={`w-5 h-5 mr-3 ${active ? 'text-blue-600' : 'text-gray-400'}`} fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><rect x="4" y="5" width="16" height="14" rx="2" stroke="currentColor" strokeWidth="2" /><path d="M8 9h8M8 13h6" stroke="currentColor" strokeWidth="2" strokeLinecap="round" /></svg>
    ),
  },
  {
    label: 'Create Task',
    icon: (active) => (
      <svg className={`w-5 h-5 mr-3 ${active ? 'text-blue-600' : 'text-gray-400'}`} fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><circle cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="2" /><path d="M12 8v8M8 12h8" stroke="currentColor" strokeWidth="2" strokeLinecap="round" /></svg>
    ),
  },
]

const user = {
  name: 'Kshitij',
  email: 'k@mail.com',
}

const getInitials = (name) => name.split(' ').map(n => n[0]).join('').toUpperCase()

const Dashboard = () => {
  const [active, setActive] = useState('Tasks')
  const [showProfile, setShowProfile] = useState(false)
  const [searchQuery, setSearchQuery] = useState('')
  const navigate = useNavigate()

  const handleLogout = () => {
    localStorage.removeItem('accessToken')
    navigate('/')
  }

  const filteredTasks = sampleTasks.filter(task =>
    task.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
    task.status.toLowerCase().includes(searchQuery.toLowerCase()) ||
    task.priority.toLowerCase().includes(searchQuery.toLowerCase())
  )

  const renderContent = () => {
    switch (active) {
      case 'Tasks':
        return (
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-bold text-gray-800">Tasks</h2>
              <div className="relative">
                <input
                  type="text"
                  placeholder="Search Task"
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="w-64 px-4 py-2 bg-red-50 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                />
                <svg
                  className="absolute right-3 top-2.5 h-5 w-5 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                  />
                </svg>
              </div>
            </div>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Priority</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Due Date</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Actions</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {filteredTasks.map((task) => (
                    <tr key={task.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{task.title}</td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full 
                          ${task.status === 'Completed' ? 'bg-green-100 text-green-800' : 
                            task.status === 'In Progress' ? 'bg-blue-100 text-blue-800' : 
                            'bg-yellow-100 text-yellow-800'}`}>
                          {task.status}
                        </span>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        <span className={`px-2 inline-flex text-xs leading-5 font-semibold rounded-full 
                          ${task.priority === 'High' ? 'bg-red-100 text-red-800' : 
                            task.priority === 'Medium' ? 'bg-yellow-100 text-yellow-800' : 
                            'bg-green-100 text-green-800'}`}>
                          {task.priority}
                        </span>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{task.dueDate}</td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                        <button className="text-blue-600 hover:text-blue-900 mr-3">Edit</button>
                        <button className="text-red-600 hover:text-red-900">Delete</button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        )
      case 'Create Task':
        return (
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-2xl font-bold text-gray-800 mb-6">Create New Task</h2>
            {/* Create task form will go here */}
          </div>
        )
      default:
        return (
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-2xl font-bold text-gray-800 mb-4">Dashboard</h2>
            <p className="text-gray-600">Welcome to your dashboard!</p>
          </div>
        )
    }
  }

  return (
    <div className="flex min-h-screen bg-gray-100">
      {/* Sidebar */}
      <aside className="w-64 bg-white h-screen flex flex-col border-r border-gray-200 fixed top-0 left-0 z-10 shadow-md">
        <div className="flex items-center h-16 px-6 font-bold text-lg text-gray-800 border-b border-gray-200">
          <span className="mr-2">
            <svg className="w-7 h-7 text-blue-600" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><rect x="3" y="3" width="18" height="18" rx="4" stroke="currentColor" strokeWidth="2" fill="#e0e7ff" /><text x="12" y="16" textAnchor="middle" fontSize="10" fill="#1e40af">1</text></svg>
          </span>
          Dashboard
        </div>
        <nav className="flex-1 py-8 px-0 space-y-1">
          {menuItems.map(item => {
            const isActive = active === item.label
            return (
              <button
                key={item.label}
                onClick={() => setActive(item.label)}
                className={`flex items-center w-full px-8 py-3 text-base font-medium transition-colors text-left focus:outline-none
                  ${isActive ? 'text-blue-600 bg-transparent border-l-4 border-blue-600' : 'text-gray-500 border-l-4 border-transparent hover:text-blue-600 hover:bg-gray-50'}`}
                style={{ background: 'none' }}
              >
                {item.icon(isActive)}
                {item.label}
              </button>
            )
          })}
        </nav>
        {/* Profile photo at bottom */}
        <div className="px-8 pb-8 mt-auto relative flex flex-col items-center">
          <div
            className="relative group"
            tabIndex={0}
            onClick={() => setShowProfile((v) => !v)}
            onBlur={() => setShowProfile(false)}
          >
            <div className="w-12 h-12 rounded-full bg-blue-200 flex items-center justify-center text-xl font-bold text-blue-800 cursor-pointer shadow-md transition-transform duration-150 hover:scale-110">
              {getInitials(user.name)}
            </div>
            {/* Popover */}
            {showProfile && (
              <div className="absolute bottom-16 left-1/2 -translate-x-1/2 mb-2 w-64 bg-white rounded-lg shadow-xl border border-gray-100 p-4 flex flex-col gap-2 animate-fade-in z-30">
                <div className="font-bold text-lg text-gray-800">{user.name}</div>
                <div className="text-sm text-gray-500 mb-2">{user.email}</div>
                <hr className="my-1" />
                <button
                  onClick={handleLogout}
                  className="text-red-600 text-base text-left hover:underline focus:outline-none bg-transparent hover:bg-transparent active:bg-transparent"
                >
                  Sign Out
                </button>
              </div>
            )}
          </div>
        </div>
      </aside>
      {/* Main Content */}
      <div className="flex-1 ml-64 p-8">
        {renderContent()}
      </div>
    </div>
  )
}

export default Dashboard 