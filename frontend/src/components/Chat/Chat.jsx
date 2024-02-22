import React, { useState } from "react";
import styles from "./Chat.module.css";

import classNames from "class-names";

import Helper from "../../assets/img/helper.svg";
import Send from "../../assets/icon/send.png";
import Close from "../../assets/icon/close.png";

const Chat = () => {
  const [isChatOpen, setIsChatOpen] = useState(false);
  const [messages, setMessages] = useState(["Hello", "pip"]);
  const [currentMessage, setCurrentMessage] = useState("");

  const handleInputChange = (event) => {
    setCurrentMessage(event.target.value);
  };

  const handleFormSubmit = (event) => {
    event.preventDefault();
    setMessages([...messages, currentMessage]);
    setCurrentMessage("");
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
              <img src={Close} alt="" />
            </button>
          </div>
          <div className={styles.chat_messages}>
            {messages.map((message, index) => (
              <div key={index}>{message}</div>
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
