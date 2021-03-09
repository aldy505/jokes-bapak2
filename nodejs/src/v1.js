const { Router } = require('pure-http');

const router = Router();

router.get('/', (req, res) => {
  // Random image with default size
  res.status(200).send('...', false, 200);
});

module.exports = router;
