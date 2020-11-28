import { Room, isRooms } from '../models/room';

const getAllRooms = async (): Promise<Room[]> => {
  const adminApiHost =
    process.env.ADMIN_API_HOST || 'http://localhost:8080/rooms';
  const response = await fetch(`http://${adminApiHost}/rooms`);

  if (!response.ok) {
    throw new Error(`${response.status} Error`);
  }

  const rooms = (await response.json()) as unknown[];

  if (!isRooms(rooms)) {
    throw new Error(`API Type Error`);
  }

  return rooms;
};

export default getAllRooms;
