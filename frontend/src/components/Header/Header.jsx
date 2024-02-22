import NavBar from "../NavBar/NavBar";
import Logo from './../../assets/img/LOGO.svg'

import styles from "./Header.module.css"

const Header = () => {
    return ( 
        <div className={styles.header_block}>
            <img src={Logo} alt="Логотип" />
            <NavBar />
        </div>
    );
}
 
export default Header;