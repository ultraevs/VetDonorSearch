import axios from "axios";

export const checkLoginForm = async (mail, password) => {
  try {
    const response = await axios.post(
      "https://vetdonor.shmyaks.ru/v1/login",
      {
        email: `${mail}`,
        password: `${password}`,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return {success: true}
  } catch (error) {
    return {success: false, error: error};
  }
};