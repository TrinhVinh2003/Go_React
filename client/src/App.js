import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import ArticleList from './components/ArticleList';
import ArticleDetail from './components/ArticleDetail';
import './App.css';
const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<ArticleList />} />
        <Route path="/articles/:id" element={<ArticleDetail />} />
      </Routes>
    </Router>
  );
};

export default App;