import React from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

export default function Dashboard() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  if (!user) {
    navigate('/login');
    return null;
  }

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold">Welcome, {user.username}</h1>
      <div className="mt-8 space-x-4">
        <button 
          onClick={() => navigate('/game')} 
          className="px-6 py-3 text-white bg-green-500 rounded font-bold"
        >
          Find Match
        </button>
        <button 
          onClick={logout} 
          className="px-6 py-3 text-white bg-red-500 rounded font-bold"
        >
          Logout
        </button>
      </div>
    </div>
  );
}
