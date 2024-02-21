import React, { useEffect, useState } from "react";
import styles from "./Register.module.css";
import Layout from "../../components/Layout/Layout";

import classNames from "class-names";

import { Link, useNavigate } from "react-router-dom";

import { createUser, createClinic } from "./http";

const Register = () => {
  const navigate = useNavigate();

  const [activeItem, setActiveItem] = useState(1);

  const [inputName, setInputName] = useState("");
  const [inputClinic, setInputClinic] = useState("");
  const [inputAdress, setInputAdress] = useState("");
  const [inputMail, setInputMail] = useState("");
  const [inputPassword, setInputPassword] = useState("");

  const clearInputs = () => {
    setInputName("");
    setInputClinic("");
    setInputAdress("");
    setInputMail("");
    setInputPassword("");
  };

  const buttonUserClick = async (name, mail, password) => {
    try {
      const response = await createUser(name, mail, password);
      if (response.success) {
        navigate("/Profile");
      } else {
        alert("Ошибка: " + response.error)
      }
      clearInputs();
    } catch (error) {
      console.error(error);
    }
  };

  const buttonClinicClick = async (name, adress, mail, password) => {
    try {
      const response = await createClinic(name, adress, mail, password);
      console.log(response)
      if (response.success) {
        navigate("/Profile");
      } else {
        alert("Ошибка: " + response.error)
      }
      clearInputs();
    } catch (error) {
      console.error(error);
    }
  };

  const handleItemClick = (item) => {
    setActiveItem(item);
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
            <p>Клиника</p>
          </div>
          <div
            className={classNames(styles.choose_item, {
              [styles.active]: activeItem === 1,
            })}
            onClick={() => handleItemClick(1)}
          >
            <p>Частный владелец</p>
          </div>
        </div>
        <div className={styles.enter_data}>
          {activeItem === 1 ? (
            <div className={styles.enter_data_name}>
              <p>ФИО</p>
              <input
                type="text"
                placeholder="Введите ФИО"
                value={inputName}
                onChange={(event) => setInputName(event.target.value)}
              />
            </div>
          ) : (
            <>
              <div className={styles.enter_data_name}>
                <p>Название</p>
                <input
                  type="text"
                  placeholder="Введите название клиники"
                  value={inputClinic}
                  onChange={(event) => setInputClinic(event.target.value)}
                />
              </div>
              <div className={styles.enter_data_name}>
                <p>Адресс</p>
                <input
                  type="text"
                  placeholder="Введите адресс клиники"
                  value={inputAdress}
                  onChange={(event) => setInputAdress(event.target.value)}
                />
              </div>
            </>
          )}
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
          <p>
            Уже есть аккаунт?{" "}
            <Link to="/auth">
              <span>Войти</span>
            </Link>
          </p>
        </div>
        <div className={styles.enter}>
          <Link>
            <button
              onClick={() => {
                if (activeItem === 1) {
                  buttonUserClick(inputName, inputMail, inputPassword);
                } else {
                  buttonClinicClick(inputName, inputAdress, inputMail, inputPassword);
                }
              }}
            >
              Зарегистрироваться
            </button>
          </Link>
        </div>
      </div>
    </Layout>
  );
};

export default Register;
