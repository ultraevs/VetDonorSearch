import React, { useState } from "react";
import Layout from "../../components/Layout/Layout";

import classNames from "class-names";

import { Link, useNavigate} from "react-router-dom";

import styles from "./Auth.module.css";

import { checkLoginForm } from "./http";

const Auth = () => {
  const navigate = useNavigate();

  const [activeItem, setActiveItem] = useState(0);

  const [inputMail, setInputMail] = useState("");
  const [inputPassword, setInputPassword] = useState("");

  const clearInputs = () => {
    setInputMail("");
    setInputPassword("");
  };

  const handleItemClick = (item) => {
    setActiveItem(item);
  };

  const buttonClick = async (mail, password) => {
    try {
      const response = await checkLoginForm(mail, password);
      console.log(response);
      if (response.success) {
        navigate("/Profile")
      } else {
        alert("Ошибка: " + response.error)
      }
      clearInputs();
    } catch (error) {
      console.error(error);
    }
  };
  
  return (
    <Layout>
      <div className={styles.main_info}>
        <div className={styles.choose}>
          <div
            className={classNames(styles.choose_item, {
              [styles.active]: activeItem === 0,
            })}
            onClick={() => handleItemClick(0)}
          >
            <p>Войти</p>
          </div>
          <Link
            to="/auth/register"
            className={classNames(styles.choose_item, {
              [styles.active]: activeItem === 1,
            })}
            onClick={() => handleItemClick(1)}
          >
            <p>Зарегистрироваться</p>
          </Link>
        </div>

        <div className={styles.enter_data}>
          <div className={styles.enter_data_mail}>
            <p>Почта</p>
            <input
              type="text"
              placeholder="Введите электронную почту"
              value={inputMail}
              onChange={(event) => setInputMail(event.target.value)}
            />
          </div>
          <div className={styles.enter_data_password}>
            <p>Пароль</p>
            <input
              type="text"
              placeholder="Введите пароль"
              value={inputPassword}
              onChange={(event) => setInputPassword(event.target.value)}
            />
          </div>
        </div>
        <div className={styles.forgot_password}>
          <p>Забыли пароль?</p>
        </div>
        <div className={styles.enter}>
          <Link>
            <button
              onClick={() => buttonClick(inputMail, inputPassword)}
            >
              Войти
            </button>
          </Link>
        </div>
      </div>
    </Layout>
  );
};

export default Auth;
