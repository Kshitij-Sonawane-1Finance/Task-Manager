import React from 'react'

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

const getInitials = (name) => name.split(' ').map(n => n[0]).join('').toUpperCase()

const Sidebar = ({ active, setActive, user, loading, showProfile, setShowProfile, handleLogout }) => {
  return (
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
      <div className="px-8 pb-8 mt-auto relative flex flex-col items-center">
        <div
          className="relative group"
          tabIndex={0}
          onClick={() => setShowProfile((v) => !v)}
        >
          <div className="w-12 h-12 rounded-full bg-blue-200 flex items-center justify-center text-xl font-bold text-blue-800 cursor-pointer shadow-md transition-transform duration-150 hover:scale-110">
            {loading ? '...' : user ? getInitials(user.name) : '?'}
          </div>
          {showProfile && (
            <div 
              className="absolute bottom-16 left-1/2 -translate-x-1/2 mb-2 w-64 bg-white rounded-lg shadow-xl border border-gray-100 p-4 flex flex-col gap-2 animate-fade-in z-30"
              onClick={(e) => e.stopPropagation()}
            >
              {loading ? (
                <div className="text-gray-500">Loading...</div>
              ) : user ? (
                <>
                  <div className="font-bold text-lg text-gray-800">{user.name}</div>
                  <div className="text-sm text-gray-500 mb-2">{user.email}</div>
                  <hr className="my-1" />
                  <button
                    type="button"
                    onClick={handleLogout}
                    className="text-red-600 text-base text-left hover:underline focus:outline-none bg-transparent hover:bg-transparent active:bg-transparent"
                  >
                    Sign Out
                  </button>
                </>
              ) : (
                <div className="text-red-500">Failed to load user data</div>
              )}
            </div>
          )}
        </div>
      </div>
    </aside>
  )
}

export default Sidebar 