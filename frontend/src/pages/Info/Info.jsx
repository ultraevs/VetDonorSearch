import MainLayout from '../../components/MainLayout/MainLayout'

import styles from './Info.module.css'

const Info = () => {
    return (
        <MainLayout>
            <div className={styles.container}>
                <div className={styles.block_white}>
                    <h3>
                        Может ли Ваш питомец спасти чью-то жизнь?
                    </h3>
                    <br />
                    <p>
                        Да! Приглашаем здоровых молодых кошек и собак стать донорами крови и помочь нуждающимся животным.<br />
                        Ежедневно к нам поступают пациенты, которым жизненно необходима донорская кровь - онкобольные, животные
                        с хроническими вирусными инфекциями, те, кто получил серьезные травмы и многие-многие другие!
                        Кроме того, донорство дает возможность получать вознаграждения.
                    </p>
                </div>
                <div className={styles.block_blue}>
                    <h3>
                        Как узнать, что Ваш друг может стать донором?
                    </h3>
                    <br />
                    <p>
                        Возраст - 1,5-7 лет<br />
                        Вес - Кошки от 3,5 кг, Собаки от 10 кг и крупнее<br />
                        Клинически здоровые<br />
                        Вакцинированы (давность вакцинации до года).<br />
                        Обработаны от экто- и эндопаразитов по графику<br />
                        Без хронических заболеваний.<br />
                        Без переливаний крови и плазмы в анамнезе<br />
                        Без онкологических заболеваний в течение жизни
                    </p>
                </div>

                <div className={styles.block_white}>
                    <h3>
                        Изменится ли как-то жизнь Вашего кота, если он станет донором крови?
                    </h3>
                    <br />
                    <p>
                        Конечно, изменится, но только в положительном ключе! <br/>
                        Теперь донор может гордиться своим званием, тем, что обеспечивает вакцинацию. Он свободен от инфекционных заболеваний спасает жизни.
                    </p>
                </div>
                <div className={styles.block_blue}>
                    <h3>
                        Есть ли какие-то побочные эффекты у донации крови?
                    </h3>
                    <br />
                    <p>
                        Сама по себе донация крови - безвредный процесс у здоровых животных. <br/>
                        Анестезиологические риски у здоровых животных также минимальны, кроме того мы проводим комплексное обследование доноров - анализы крови и кардиологическое обследование.
                    </p>
                </div>

                <div className={styles.block_white}>
                    <h3>
                        Нужен ли какой-то особенный уход донору после сдачи крови?
                    </h3>
                    <br />
                    <p>
                        Особенный уход не нужен, нужно просто покормить донора по возвращении домой и обеспечить неограниченный доступ к воде.
                    </p>
                </div>
                <div className={styles.block_blue}>
                    <h3>
                        Как часто животное может быть донором?
                    </h3>
                    <br />
                    <p>
                        Оптимально сдавать кровь не чаще, чем 1 раз в 3 месяца, чтобы донор мог полностью восстановиться.
                    </p>
                </div>

                <div className={styles.block_white}>
                    <h3>
                        Я решил, что мой питомец станет донором. Что дальше?
                    </h3>
                    <br />
                    <p>
                        Вам необходимо позвонить в клинику, если ближайшая дата «дня донора» известна и вам удобен филиал и время, то вас с питомцем внесут в список доноров, если дата еще не определена, запишут ваши контакты и удобный для вас филиал (Митино, Пресня, Строгино, Крылатское). Вам позвонят и пригласят на донацию, когда она будет проводиться в озвученном вами филиале.
                    </p>
                </div>
                <div className={styles.block_blue}>
                    <h3>
                        Вознаграждение
                    </h3>
                    <br />
                    <p>
                        Такой важный и нелегкий поступок для вас и вашего питомца не останутся без награды.Бесплатный осмотр врача и анализы крови, необходимые для контроля здоровья животного, шанс спасти жизнь другому животному, денежное вознаграждение.
                    </p>
                </div>

            </div>
        </MainLayout>
    );
}

export default Info;