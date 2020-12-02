import camelcaseKeys from 'camelcase-keys';
import { Room, isRoom } from '../models/room';

const createNewRoom = async (newRoom: {name: string, description: string, hostName: string}): Promise<Room> => {
  const adminApiHost = process.env.ADMIN_API_HOST || 'localhost:8080';
  const response = await fetch(`http://${adminApiHost}/rooms`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(newRoom),
  });

  if (!response.ok) {
    throw new Error(`${response.status} Error`);
  }

  const room = camelcaseKeys(await response.json()) as unknown[];

  if (!isRoom(room)) {
    throw new Error(`API Type Error`);
  }

  return room;
};

export default createNewRoom;
