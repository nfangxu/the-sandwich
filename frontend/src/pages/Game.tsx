import React, { useEffect, useState, useRef } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

interface GameState {
  MatchID: string;
  Status: string;
  Players: string[];
}

export default function Game() {
  const { user, token } = useAuth();
  const navigate = useNavigate();
  const [gameState, setGameState] = useState<GameState | null>(null);
  const [statusText, setStatusText] = useState('Connecting...');
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!user || !token) {
      navigate('/login');
      return;
    }

    const socket = new WebSocket(`ws://localhost:8080/ws?token=${token}`);
    
    socket.onopen = () => {
      setStatusText('Connected. Waiting for match...');
      // Send join queue message (assuming the backend hub expects JSON)
      socket.send(JSON.stringify({ type: 'JOIN_QUEUE' }));
    };

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        if (data.type === 'STATE_UPDATE') {
          setGameState(data.state);
          setStatusText(`Match ${data.state.MatchID} - Status: ${data.state.Status}`);
        }
      } catch (e) {
        console.log('Received raw message:', event.data);
      }
    };

    socket.onclose = () => {
      setStatusText('Disconnected.');
    };

    ws.current = socket;

    return () => {
      socket.close();
    };
  }, [user, token, navigate]);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-green-800 text-white">
      <h1 className="text-2xl font-bold mb-4">The Sandwich Table</h1>
      <p className="mb-8 text-xl">{statusText}</p>
      
      {gameState && (
        <div className="p-8 bg-green-700 rounded-lg shadow-lg">
          <h2 className="text-xl mb-4">Players in match:</h2>
          <ul>
            {gameState.Players?.map((p, i) => (
              <li key={i} className="mb-2">Player: {p}</li>
            ))}
          </ul>
        </div>
      )}
      
      <button 
        className="mt-8 px-4 py-2 bg-red-600 rounded font-bold hover:bg-red-700"
        onClick={() => navigate('/')}
      >
        Leave Table
      </button>
    </div>
  );
}
