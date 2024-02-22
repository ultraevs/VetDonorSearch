import React from 'react'
import Title from './components/Title/Title'
import MessageList from './components/MessageList/MessageList'
import SendMessageForm from './components/SendMessageForm/SendMessageForm'

const Chat = () => {
  return (
    <div className='chat'>
        <Title />
        <MessageList />
        <SendMessageForm />
    </div>
  )
}

export default Chat