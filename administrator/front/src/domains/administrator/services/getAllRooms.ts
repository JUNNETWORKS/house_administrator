import camelcaseKeys from 'camelcase-keys';
import { Room, isRooms } from '../models/room';

const getAllRooms = async (): Promise<Room[]> => {
  const adminApiHost = process.env.ADMIN_API_HOST || 'localhost:8080';
  const response = await fetch(`http://${adminApiHost}/rooms`);

  if (!response.ok) {
    throw new Error(`${response.status} Error`);
  }

  const rooms = camelcaseKeys(await response.json()) as unknown[];

  if (!isRooms(rooms)) {
    throw new Error(`API Type Error`);
  }

  return rooms;
};

export default getAllRooms;
