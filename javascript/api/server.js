const express = require('express');
const app = express();

app.use(express.json()); // Middleware to parse JSON bodies

app.get('/', (req, res) => {
  res.status(200).json({ message: '🎉 Your Express server is running' });
});

app.post('/', (req, res) => {
  res.status(201).json({ message: '✅ POST request received successfully' });
});

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`🚀 Server is listening on port ${PORT}`);
});
