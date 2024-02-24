import React, { useState } from "react";
import Layout from "../../components/Layout/Layout";

import classNames from "class-names";

import { Link, useNavigate } from "react-router-dom";

import styles from "./Forgot.module.css";

import { checkForgotForm } from "./http";

import Cookies from "js-cookie";

const Forgot = () => {
    const navigate = useNavigate();

    const [activeItem, setActiveItem] = useState(0);

    const [inputMail, setInputMail] = useState("");

    const clearInputs = () => {
        setInputMail("");
    };

    const handleItemClick = (item) => {
        setActiveItem(item);
    };

    const buttonClick = async (mail) => {
        try {
            const response = await checkForgotForm(mail);
            if (response.success) {
                console.log(response)
                window.alert("На вашу почту отправлена ссылка для восстановления")
            } else {
                alert("Ошибка: " + response.error);
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
                    <Link
                        to="/auth"
                        className={classNames(styles.choose_item,)}
                        onClick={() => handleItemClick(0)}
                    >
                        <p>Войти</p>
                    </Link>
                    <Link
                        to="/auth/register"
                        className={classNames(styles.choose_item,)}
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
                </div>
                <div className={styles.enter}>
                    <Link>
                        <button
                            className="button"
                            onClick={() => buttonClick(inputMail)}
                        >
                            Восстановить
                        </button>
                    </Link>
                </div>
            </div>
        </Layout>
    );
};

export default Forgot;