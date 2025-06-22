import React, { useState, useEffect, Fragment } from 'react'
import { Listbox, Transition } from '@headlessui/react'
import { CheckIcon, SelectorIcon } from '@heroicons/react/solid'
import axios from 'axios'
import TaskDetailsModal from './TaskDetailsModal'

const statusOptions = [
  { name: 'All' },
  { name: 'New' },
  { name: 'Active' },
  { name: 'Completed' },
  { name: 'Deferred' },
]

const priorityOptions = [
  { name: 'All' },
  { name: 'High' },
  { name: 'Medium' },
  { name: 'Low' },
]

const sortOptions = [
  { name: '🕐 Earliest Start' },
  { name: '🕓 Latest Start' },
  { name: '📅 Due Soon' },
  { name: '🗓️ Due Later' },
];

const TaskTable = ({ 
  tasks, 
  searchQuery, 
  onEditSuccess, 
  onView, 
  onDeleteSuccess, 
  page, 
  setPage, 
  count,
  onFiltersChange 
}) => {
  const [selectedTask, setSelectedTask] = useState(null)
  const [popup, setPopup] = useState({ show: false, message: '', success: false })
  const [searchValue, setSearchValue] = useState(searchQuery)
  const [searchResults, setSearchResults] = useState(null)
  const [searchLoading, setSearchLoading] = useState(false)
  const [selectedStatus, setSelectedStatus] = useState(statusOptions[0])
  const [selectedPriority, setSelectedPriority] = useState(priorityOptions[0])
  const [selectedSort, setSelectedSort] = useState(sortOptions[2])
  const isInitialMount = React.useRef(true);

  const totalPages = Math.ceil(count / 5) || 1;

  // Function to get sortBy and orderBy from selected sort option
  const getSortParams = (sortOption) => {
    switch (sortOption.name) {
      case '🕐 Earliest Start':
        return { sortBy: 'start_date', orderBy: 'asc' }
      case '🕓 Latest Start':
        return { sortBy: 'start_date', orderBy: 'desc' }
      case '📅 Due Soon':
        return { sortBy: 'end_date', orderBy: 'desc' }
      case '🗓️ Due Later':
        return { sortBy: 'end_date', orderBy: 'asc' }
      default:
        return { sortBy: 'end_date', orderBy: 'desc' }
    }
  }

  // Function to send unified API request
  const sendUnifiedRequest = async () => {
    try {
      const token = localStorage.getItem('accessToken')
      const params = new URLSearchParams({
        page: page,
        limit: 5
      })

      // Add search data if exists
      if (searchValue && searchValue.trim() !== '') {
        params.append('searchData', searchValue)
      }

      // Add status if not "All"
      if (selectedStatus.name !== 'All') {
        params.append('status', selectedStatus.name)
      }

      // Add priority if not "All"
      if (selectedPriority.name !== 'All') {
        params.append('priority', selectedPriority.name)
      }

      // Add sort parameters
      const sortParams = getSortParams(selectedSort)
      params.append('sortBy', sortParams.sortBy)
      params.append('orderBy', sortParams.orderBy)

      const response = await axios.get(`http://localhost:3000/api/task?${params.toString()}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      if (response.data) {
        // Call the parent callback to update the tasks
        if (onFiltersChange) {
          onFiltersChange(response.data.tasks, response.data.count)
        }
      }
    } catch (error) {
      console.error('Failed to fetch tasks:', error)
    }
  }

  // Debounced search effect
  useEffect(() => {
    if (isInitialMount.current) {
      isInitialMount.current = false;
      return;
    }

    const delayDebounce = setTimeout(() => {
      sendUnifiedRequest()
    }, 500)

    return () => clearTimeout(delayDebounce)
  }, [searchValue])

  // Effect for status, priority, and sort changes
  useEffect(() => {
    sendUnifiedRequest()
  }, [selectedStatus, selectedPriority, selectedSort])

  // Effect for page changes
  useEffect(() => {
    sendUnifiedRequest()
  }, [page])

  // Update search value when searchQuery prop changes
  useEffect(() => {
    setSearchValue(searchQuery)
  }, [searchQuery])

  const filteredTasks = searchResults !== null ? searchResults : tasks.filter(task =>
    task.title.toLowerCase().includes(searchValue?.toLowerCase() || '') ||
    task.status.toLowerCase().includes(searchValue?.toLowerCase() || '') ||
    task.priority.toLowerCase().includes(searchValue?.toLowerCase() || '')
  )

  const handleViewTask = async (taskId) => {
    try {
      const token = localStorage.getItem('accessToken')
      const response = await axios.get(`http://localhost:3000/api/task/${taskId}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })

      if (response.data) {
        setSelectedTask(response.data)
      }
    } catch (error) {
      console.error('Error fetching task details:', error)
      alert('Failed to fetch task details')
    }
  }

  const handleDelete = async (taskId) => {
    if (!window.confirm('Are you sure you want to delete this task?')) return
    try {
      const token = localStorage.getItem('accessToken')
      const response = await axios.delete(`http://localhost:3000/api/task/${taskId}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      })
      if (response.status === 200) {
        setPopup({ show: true, message: response.data.message, success: true })
        if (onDeleteSuccess) onDeleteSuccess()
        setTimeout(() => setPopup({ show: false, message: '', success: false }), 2000)
      }
    } catch (error) {
      setPopup({ show: true, message: 'Failed to delete task.', success: false })
      setTimeout(() => setPopup({ show: false, message: '', success: false }), 2000)
    }
  }

  return (
    <div className="bg-white rounded-lg shadow-md p-6">
      <div className="mb-4 flex items-center">
        <h2 className="text-2xl font-bold text-gray-800">Tasks</h2>
      </div>
      {popup.show && (
        <div className={`fixed top-6 left-1/2 transform -translate-x-1/2 z-50 flex items-center gap-2 px-6 py-3 rounded shadow-lg text-white ${popup.success ? 'bg-green-600' : 'bg-red-600'}`}>
          <span>{popup.message}</span>
        </div>
      )}
      <div className="flex items-end gap-4 mb-6 justify-between">
        <div className="flex items-end gap-4">
          {/* Custom Status Filter */}
          <div className="flex flex-col items-start w-40">
            <label className="mb-1 ml-1 text-xs font-medium text-gray-600">Status</label>
            <div className="w-full relative">
              <Listbox value={selectedStatus} onChange={setSelectedStatus}>
                <div className="relative">
                  <Listbox.Button className="relative w-full cursor-pointer rounded-md bg-white py-1.5 pl-9 pr-8 text-left border border-gray-200 text-sm text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <span className="absolute left-2 inset-y-0 flex items-center">
                      {/* Status Icon */}
                      <svg className="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24">
                        <circle cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="2" fill="none" />
                        <path strokeLinecap="round" strokeLinejoin="round" d="M8 12l2 2 4-4" />
                      </svg>
                    </span>
                    <span className="block truncate pl-5">{selectedStatus.name}</span>
                    <span className="pointer-events-none absolute inset-y-0 right-2 flex items-center">▼</span>
                  </Listbox.Button>
                  <Transition as={Fragment} leave="transition ease-in duration-100" leaveFrom="opacity-100" leaveTo="opacity-0">
                    <Listbox.Options className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                      {statusOptions.map((option, idx) => (
                        <Listbox.Option
                          key={option.name}
                          className={({ active }) =>
                            `relative cursor-pointer select-none py-2 pl-10 pr-4 ${active ? 'bg-blue-100 text-blue-900' : 'text-gray-900'}`
                          }
                          value={option}
                        >
                          {({ selected }) => (
                            <>
                              <span className={`block truncate ${selected ? 'font-medium' : 'font-normal'}`}>{option.name}</span>
                              {selected ? (
                                <span className="absolute inset-y-0 left-2 flex items-center text-blue-600">✓</span>
                              ) : null}
                            </>
                          )}
                        </Listbox.Option>
                      ))}
                    </Listbox.Options>
                  </Transition>
                </div>
              </Listbox>
            </div>
          </div>
          {/* Custom Priority Filter */}
          <div className="flex flex-col items-start w-40">
            <label className="mb-1 ml-1 text-xs font-medium text-gray-600">Priority</label>
            <div className="w-full relative">
              <Listbox value={selectedPriority} onChange={setSelectedPriority}>
                <div className="relative">
                  <Listbox.Button className="relative w-full cursor-pointer rounded-md bg-white py-1.5 pl-9 pr-8 text-left border border-gray-200 text-sm text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <span className="absolute left-2 inset-y-0 flex items-center">
                      {/* Priority Icon */}
                      <svg className="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24">
                        <polygon points="12,2 15,8 22,9 17,14 18,21 12,18 6,21 7,14 2,9 9,8" />
                      </svg>
                    </span>
                    <span className="block truncate pl-5">{selectedPriority.name}</span>
                    <span className="pointer-events-none absolute inset-y-0 right-2 flex items-center">▼</span>
                  </Listbox.Button>
                  <Transition as={Fragment} leave="transition ease-in duration-100" leaveFrom="opacity-100" leaveTo="opacity-0">
                    <Listbox.Options className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                      {priorityOptions.map((option, idx) => (
                        <Listbox.Option
                          key={option.name}
                          className={({ active }) =>
                            `relative cursor-pointer select-none py-2 pl-10 pr-4 ${active ? 'bg-blue-100 text-blue-900' : 'text-gray-900'}`
                          }
                          value={option}
                        >
                          {({ selected }) => (
                            <>
                              <span className={`block truncate ${selected ? 'font-medium' : 'font-normal'}`}>{option.name}</span>
                              {selected ? (
                                <span className="absolute inset-y-0 left-2 flex items-center text-blue-600">✓</span>
                              ) : null}
                            </>
                          )}
                        </Listbox.Option>
                      ))}
                    </Listbox.Options>
                  </Transition>
                </div>
              </Listbox>
            </div>
          </div>
          {/* Custom Sort By */}
          <div className="flex flex-col items-start w-44">
            <label className="mb-1 ml-1 text-xs font-medium text-gray-600">Sort By</label>
            <div className="w-full relative">
              <Listbox value={selectedSort} onChange={setSelectedSort}>
                <div className="relative">
                  <Listbox.Button className="relative w-full cursor-pointer rounded-md bg-white py-1.5 pl-9 pr-8 text-left border border-gray-200 text-sm text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
                    <span className="absolute left-2 inset-y-0 flex items-center">
                      {/* Sort Icon */}
                      <svg className="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M3 7h18M3 12h12M3 17h6" />
                      </svg>
                    </span>
                    <span className="block truncate pl-5">{selectedSort.name}</span>
                    <span className="pointer-events-none absolute inset-y-0 right-2 flex items-center">▼</span>
                  </Listbox.Button>
                  <Transition as={Fragment} leave="transition ease-in duration-100" leaveFrom="opacity-100" leaveTo="opacity-0">
                    <Listbox.Options className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                      {sortOptions.map((option, idx) => (
                        <Listbox.Option
                          key={option.name}
                          className={({ active }) =>
                            `relative cursor-pointer select-none py-2 pl-10 pr-4 ${active ? 'bg-blue-100 text-blue-900' : 'text-gray-900'}`
                          }
                          value={option}
                        >
                          {({ selected }) => (
                            <>
                              <span className={`block truncate ${selected ? 'font-medium' : 'font-normal'}`}>{option.name}</span>
                              {selected ? (
                                <span className="absolute inset-y-0 left-2 flex items-center text-blue-600">✓</span>
                              ) : null}
                            </>
                          )}
                        </Listbox.Option>
                      ))}
                    </Listbox.Options>
                  </Transition>
                </div>
              </Listbox>
            </div>
          </div>
        </div>
        {/* Search */}
        <div className="flex flex-col items-start w-64 ml-auto">
          <label className="mb-1 ml-1 text-xs font-medium text-gray-600">Search</label>
          <div className="relative w-full">
            <input
              type="text"
              placeholder="Search Task"
              value={searchValue}
              onChange={(e) => setSearchValue(e.target.value)}
              className="w-full px-4 py-1.5 bg-red-50 border border-gray-200 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent text-gray-900 text-sm"
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
            {searchLoading && <span className="absolute right-10 top-2.5 text-xs text-gray-400">Searching...</span>}
          </div>
        </div>
      </div>
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Title</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Priority</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Start Date</th>
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
                      task.status === 'Active' ? 'bg-blue-100 text-blue-800' : 
                      task.status === 'New' ? 'bg-yellow-100 text-yellow-800' :
                      'bg-gray-100 text-gray-800'}`}>
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
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{task.start_date}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{task.end_date}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <button 
                    onClick={() => handleViewTask(task.id)}
                    className="px-3 py-1 bg-blue-50 text-blue-700 rounded-md hover:bg-blue-100 transition-colors mr-2"
                  >
                    View
                  </button>
                  <button 
                    className="px-3 py-1 bg-red-50 text-red-700 rounded-md hover:bg-red-100 transition-colors"
                    onClick={() => handleDelete(task.id)}
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      {/* Pagination Controls */}
      <div className="flex justify-end items-center mt-4 gap-2">
        <button
          className="px-3 py-1 rounded bg-gray-200 text-gray-700 hover:bg-gray-300 disabled:opacity-50"
          onClick={() => setPage(page - 1)}
          disabled={page === 1}
        >
          Prev
        </button>
        {Array.from({ length: totalPages }, (_, i) => (
          <button
            key={i + 1}
            className={`px-3 py-1 rounded border ${page === i + 1 ? 'bg-blue-600 text-white border-blue-600' : 'bg-gray-100 text-gray-700 border-gray-200 hover:bg-gray-200'}`}
            onClick={() => setPage(i + 1)}
          >
            {i + 1}
          </button>
        ))}
        <button
          className="px-3 py-1 rounded bg-gray-200 text-gray-700 hover:bg-gray-300 disabled:opacity-50"
          onClick={() => setPage(page + 1)}
          disabled={page === totalPages}
        >
          Next
        </button>
      </div>
      {selectedTask && (
        <TaskDetailsModal
          task={selectedTask}
          onClose={() => setSelectedTask(null)}
          onEditSuccess={sendUnifiedRequest}
        />
      )}
    </div>
  )
}

export default TaskTable 