import React from 'react';
import './LoginButtonStyle.css'

import { oauthRequest } from '../utils/oauthRequest'

function loginButton() {
    return <button onClick={() => oauthRequest()} >Logar</button>
}

export default loginButton