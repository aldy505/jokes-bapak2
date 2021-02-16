const polka = require('polka')
const cors = require('cors')
const rateLimit = require('express-rate-limit')


const app = polka()
const limiter = rateLimit({
    windowMs: 15 * 60 * 1000, // 15 minutes
    max: 100 // limit each IP to 100 requests per windowMs
  });
app.use(cors())
app.use(limiter())
