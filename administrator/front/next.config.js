require('dotenv').config();
module.exports = {
  env: {
    // Reference a variable that was defined in the .env file and make it available at Build Time
    ADMIN_API_HOST: process.env.ADMIN_API_HOST,
  },
}
