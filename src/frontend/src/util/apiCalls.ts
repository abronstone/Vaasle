import axios from 'axios';

export const engineApi = async (): Promise<string> => {
  try {
    const res = await axios.get("http://localhost:5001/");

    if (res.status !== 200) {
        throw new Error(`HTTP error! Status: ${res.status}`);
    }

    const data = await res.data;
    return JSON.stringify(data);
  } catch (e) {
    throw new Error("Fetch failed!!");
  }
};