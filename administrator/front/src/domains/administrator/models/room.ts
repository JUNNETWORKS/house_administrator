export type Room = {
  id: number;
  name: string;
  description: string;
  hostName: string;
  ownerId?: number;
  createdAt?: Date;
  updatedAt?: Date;
};

const isRoom = (arg: unknown): arg is Room => {
  const u = arg as Room;

  return (
    typeof u?.id === 'number' &&
    typeof u?.name === 'string' &&
    typeof u?.description === 'string' &&
    typeof u?.hostName === 'string' &&
    (typeof u?.ownerId === 'number' || typeof u?.ownerId === 'undefined')
  );
};

const isRooms = (args: unknown[]): args is Room[] =>
  !args.some((arg) => !isRoom(arg));

export { isRoom, isRooms };
