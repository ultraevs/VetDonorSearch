import React, { useState, useEffect, useRef } from "react";
import styles from "./Chat.module.css";

import Helper from "../../assets/img/helper.svg";
import Send from "../../assets/icon/send.png";

const Chat = () => {
  const [isChatOpen, setIsChatOpen] = useState(false);
  const [messages, setMessages] = useState(["Привет!", "Напиши мне что-нибудь."]);
  const [currentMessage, setCurrentMessage] = useState("");

  const handleInputChange = (event) => {
    setCurrentMessage(event.target.value);
  };

  const handleFormSubmit = (event) => {
    event.preventDefault();
    setMessages([...messages, currentMessage]);
    setCurrentMessage("");
  };

  const chatContainerRef = useRef(null);

  useEffect(() => {
    if (chatContainerRef.current) {
      scrollToBottom();
    }
  }, [messages]);

  const scrollToBottom = () => {
    chatContainerRef.current.scrollTop = chatContainerRef.current.scrollHeight;
  };

  const handleOpenChat = () => {
    setIsChatOpen(true);
  };

  const handleCloseChat = () => {
    setIsChatOpen(false);
  };

  const handleOutsideClick = (event) => {
    if (
      !event.target.closest(".chat") &&
      !event.target.closest(".chat-outside")
    ) {
      setIsChatOpen(false);
    }
  };

  return (
    <div
      className={
        isChatOpen ? `${styles.chat} ${styles.open}` : `${styles.chat}`
      }
    >
      {isChatOpen !== true && (
        <button className={styles.chat_open_button} onClick={handleOpenChat}>
          <img src={Helper} alt="" />
        </button>
      )}
      {isChatOpen && (
        <div className={styles.chat_container}>
          <div className={styles.chat_header}>
            <p>Умный помощник</p>
            <button
              className={styles.chat_close_button}
              onClick={handleCloseChat}
            >
              <i class="fa-solid fa-xmark"></i>
            </button>
          </div>
          <div className={styles.chat_messages} ref={chatContainerRef}>
            {messages.map((message, index) => (
              <div className={styles.chat_message} key={index}>
                <p>{message}</p>
              </div>
            ))}
          </div>
          <form onSubmit={handleFormSubmit}>
            <input
              type="text"
              placeholder="Введите сообщение..."
              value={currentMessage}
              onChange={handleInputChange}
            />
            <button type="submit">
              <img src={Send} alt="" />
            </button>
          </form>
        </div>
      )}
      {isChatOpen && (
        <div className={styles.chat_outside} onClick={handleOutsideClick}></div>
      )}
    </div>
  );
};

export default Chat;
