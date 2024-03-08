const express = require('express');
const readme = require('readmeio');
const dotenv = require('dotenv');
const winston = require('winston');

dotenv.config(); // Load environment variables from .env file

const app = express();

// Your ReadMe secret
const secret = process.env.README_API_KEY;

// Configure Winston logger
const logger = winston.createLogger({
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.simple()
  ),
  transports: [
    new winston.transports.Console()
  ]
});

// Middleware to log incoming requests
app.use((req, res, next) => {
  logger.info(`${req.method} ${req.url}`);
  next();
});

app.post('/webhook', express.json({ type: 'application/json' }), async (req, res) => {
    // Verify the request is legitimate and came from ReadMe.
    const signature = req.headers['readme-signature'];

    try {
	readme.verifyWebhook(req.body, signature, secret);
    } catch (e) {
	// Handle invalid requests
	logger.error(`Error: ${e.message}`);
	return res.status(401).json({ error: e.message });
    }

    // Fetch the user from the database and return their data for use with OpenAPI variables.
    // const user = await db.find({ email: req.body.email })
    logger.info('Webhook processed successfully');
    return res.json({
	// Add custom data to return in your webhook call here.
	email: req.body.email,
    });
});

const server = app.listen(8000, '0.0.0.0', () => {
    console.log('Webhook app listening at http://%s:%s', server.address().address, server.address().port);
});
