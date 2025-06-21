import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import TaskTable from '../components/TaskTable'
import TaskForm from '../components/TaskForm'
import Sidebar from '../components/Sidebar'
import TaskDetailsModal from '../components/TaskDetailsModal'

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

const getInitials = (name) => name.split(' ').map(n => n[0]).join('').toUpperCase()

const Dashboard = () => {
  const [active, setActive] = useState('Tasks')
  const [showProfile, setShowProfile] = useState(false)
  const [searchQuery, setSearchQuery] = useState('')
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)
  const [tasks, setTasks] = useState([])
  const [tasksLoading, setTasksLoading] = useState(true)
  const [successMessage, setSuccessMessage] = useState('')
  const [editingTask, setEditingTask] = useState(null)
  const [selectedTaskId, setSelectedTaskId] = useState(null)
  const [page, setPage] = useState(1)
  const [count, setCount] = useState(0)
  const navigate = useNavigate()

  const fetchTasks = async () => {
    console.log('Fetching tasks...');
    try {
      const token = localStorage.getItem('accessToken')
      if (!token) {
        navigate('/')
        return
      }
      // Use current page
      const response = await axios.get(`http://localhost:3000/api/task?page=${page}&limit=5&sortBy=end_date&orderBy=desc`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      if (response.data) {
        setTasks(response.data.tasks)
        setCount(response.data.count)
      }
    } catch (error) {
      console.error('Error fetching tasks:', error)
      if (error.response?.status === 401) {
        localStorage.removeItem('accessToken')
        navigate('/')
      }
    } finally {
      setTasksLoading(false)
    }
  }

  useEffect(() => {
    if (active === 'Tasks') {
      fetchTasks()
    }
  }, [active, navigate, page])

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const token = localStorage.getItem('accessToken')
        if (!token) {
          navigate('/')
          return
        }

        const response = await axios.get('http://localhost:3000/api/user/fetchUser', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })

        if (response.data) {
          setUser(response.data)
        } else {
          console.error('Failed to fetch user data')
        }
      } catch (error) {
        console.error('Error fetching user data:', error)
        if (error.response?.status === 401) {
          localStorage.removeItem('accessToken')
          navigate('/')
        }
      } finally {
        setLoading(false)
      }
    }

    fetchUserData()
  }, [navigate])

  const fetchTaskById = async (taskId) => {
    try {
      const token = localStorage.getItem('accessToken')
      const response = await axios.get(`http://localhost:3000/api/task/${taskId}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      if (response.data) {
        return response.data
      }
    } catch (error) {
      console.error('Error fetching task details:', error)
      alert('Failed to fetch task details')
    }
    return null
  }

  const handleEdit = async (taskId) => {
    if (!taskId) return
    const task = await fetchTaskById(taskId)
    if (task) {
      setEditingTask(task)
      setSelectedTaskId(null)
      setActive('Create Task')
    }
  }

  // Show details modal for a task
  const handleViewTask = (taskId) => {
    setSelectedTaskId(taskId)
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    const formData = new FormData(e.target)
    const taskData = {
      title: formData.get('title'),
      description: formData.get('description'),
      start_date: formData.get('start_date'),
      end_date: formData.get('end_date'),
      priority: formData.get('priority'),
      status: formData.get('status')
    }

    try {
      const token = localStorage.getItem('accessToken')
      let response
      let updatedTaskId = null

      if (editingTask) {
        response = await axios.put(`http://localhost:3000/api/task/${editingTask.id}`, taskData, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        console.log("PUT response:", response);
        updatedTaskId = editingTask.id
      } else {
        response = await axios.post('http://localhost:3000/api/task/', taskData, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        })
        console.log("POST response:", response);
        updatedTaskId = response.data?.id
      }

      if (response.status === 201 || response.status === 200) {
        console.log("Reloading...");
        fetchTasks()
        setSuccessMessage(response.data.message || (editingTask ? 'Task updated successfully!' : 'Task created successfully!'))
        e.target.reset()
        setEditingTask(null)
        setSelectedTaskId(updatedTaskId)
      }
    } catch (error) {
      console.log("Error in handleSubmit:", error);
      console.error('Error saving task:', error)
      setSuccessMessage('Failed to save task. Please try again.')
      setTimeout(() => {
        setSuccessMessage('')
      }, 3000)
    }
  }

  const handleLogout = async () => {
    try {
      localStorage.removeItem('accessToken')
      setShowProfile(false)
      navigate('/', { replace: true })
    } catch (error) {
      console.error('Error during logout:', error)
    }
  }

  const handleTaskUpdated = () => {
    fetchTasks()
  }

  const handleFiltersChange = (newTasks, newCount) => {
    setTasks(newTasks)
    setCount(newCount)
  }

  const renderContent = () => {
    switch (active) {
      case 'Tasks':
        return (
          <TaskTable 
            tasks={tasks}
            searchQuery={searchQuery}
            onEdit={handleEdit}
            onView={handleViewTask}
            onDeleteSuccess={fetchTasks}
            page={page}
            setPage={setPage}
            count={count}
            onTaskUpdated={handleTaskUpdated}
            onFiltersChange={handleFiltersChange}
          />
        )
      case 'Create Task':
        return (
          <div>
            {successMessage && (
              <div className="mb-4 p-4 bg-green-100 border border-green-400 text-green-700 rounded-md">
                {successMessage}
              </div>
            )}
            <TaskForm
              editingTask={editingTask}
              onSubmit={handleSubmit}
              onCancel={() => {
                setEditingTask(null)
                setActive('Tasks')
              }}
            />
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

  // Find the selected task object for the modal
  const selectedTask = tasks.find(t => t.id === selectedTaskId)

  useEffect(() => {
    if (selectedTaskId === null) {
      fetchTasks();
    }
  }, [selectedTaskId]);

  return (
    <div className="flex min-h-screen bg-gray-100">
      <Sidebar
        active={active}
        setActive={setActive}
        user={user}
        loading={loading}
        showProfile={showProfile}
        setShowProfile={setShowProfile}
        handleLogout={handleLogout}
      />
      <div className="flex-1 ml-64 p-8">
        {renderContent()}
        {selectedTaskId && selectedTask && (
          <TaskDetailsModal
            task={selectedTask}
            onClose={() => {
              console.log('Setting selectedTaskId to null');
              setSelectedTaskId(null);
            }}
            onEdit={handleEdit}
          />
        )}
      </div>
    </div>
  )
}

export default Dashboard 