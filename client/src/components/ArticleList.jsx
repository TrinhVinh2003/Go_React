import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import './ArticleList.css'; // Import CSS

const ArticleList = () => {
  const [articles, setArticles] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8000/api/vnexpress')
      .then(response => {
        setArticles(response.data.blog_records);
      })
      .catch(error => {
        console.error("Error!", error);
      });
  }, []);

  return (
    <div className="article-list">
      <h1>vnexpress</h1>
      <ul>
        {articles.map(article => (
          <li key={article.id}>
            <Link to={`/articles/${article.id}`}>
            <div className="img1">
              <img src={article.image} alt={article.title} />
            </div>
              
             
            </Link>
              <div className="title">
                  <Link to={`/articles/${article.id}`}>
                    <h2>{article.title}</h2>
                  </Link>
                  <p>{article.description}</p>
              </div>
              
    
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ArticleList;