import React from 'react'
import { Routes, Route, Navigate, useLocation } from 'react-router-dom'
import Navbar from './components/Navbar.jsx'
import Login from './pages/Login.jsx'
import Dashboard from './pages/Dashboard.jsx'

// PrivateRoute: Only accessible if logged in
function PrivateRoute({ children }) {
  const token = localStorage.getItem('accessToken')
  const location = useLocation()
  return token ? children : <Navigate to="/login" state={{ from: location }} replace />
}

// PublicRoute: Only accessible if NOT logged in
function PublicRoute({ children }) {
  const token = localStorage.getItem('accessToken')
  return !token ? children : <Navigate to="/dashboard" replace />
}

function App() {
  const location = useLocation()
  const isDashboard = location.pathname.startsWith('/dashboard')
  return (
    <div className="min-h-screen bg-gray-100 flex flex-col">
      {!isDashboard && <Navbar />}
      <main className="flex-1 flex flex-col items-center justify-center px-4 py-8 w-full">
        <Routes>
          <Route path="/" element={
            <PublicRoute>
              <div className="text-center w-full">
                <h1 className="text-4xl font-bold text-gray-900 mb-4">
                  Welcome to Task Manager
                </h1>
                <p className="text-lg text-gray-600">
                  Organize your tasks efficiently
                </p>
              </div>
            </PublicRoute>
          } />
          <Route path="/login" element={
            <PublicRoute>
              <Login />
            </PublicRoute>
          } />
          <Route path="/register" element={
            <PublicRoute>
              <div>Register Page</div>
            </PublicRoute>
          } />
          <Route path="/dashboard" element={
            <PrivateRoute>
              <Dashboard />
            </PrivateRoute>
          } />
        </Routes>
      </main>
    </div>
  )
}

export default App