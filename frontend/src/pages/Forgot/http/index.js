import axios from "axios";

export const checkForgotForm = async (mail, password) => {
  try {
    const response = await axios.post(
      "https://vetdonor.shmyaks.ru/v1/forgot",
      {
        email: `${mail}`,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return {success: true, data: response.data.token}
  } catch (error) {
    return {success: false, error: error};
  }
};