import React from 'react';
import ReactDOM from 'react-dom';
import { Auth0Provider } from '@auth0/auth0-react';
import './index.css';
import App from './App';

ReactDOM.render(
  <Auth0Provider
    domain="dev-aj4r5fkdxle6ccqk.us.auth0.com"
    clientId="G3zHxybsKblaoLOQ3xYvI6MtlqrjIPAh"
    authorizationParams={{
      redirect_uri: window.location.origin
    }}
  >
  <React.StrictMode>
    <App />
  </React.StrictMode>
  </Auth0Provider>,
  document.getElementById('root')
)