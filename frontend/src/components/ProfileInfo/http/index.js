import axios from "axios";

export const sendInfoUser = async (city, email, name, path, phone, telegram) => {
  try {
    const response = await axios.post(
      "https://vetdonor.shmyaks.ru/v1/create_other_info",
      {
        city: `${city}`,
        email: `${email}`,
        name: `${name}`,
        path: `${path}`,
        phone: `${phone}`,
        telegram: `${telegram}`,
      },
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return { success: true};
  } catch (error) {
    return { success: false, error: error };
  }
};
