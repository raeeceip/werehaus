import request from 'supertest';
import { app } from '../src/index'; // Adjust the import path as needed

describe('Server Routes', () => {
    // Test for GET /
    test('GET / should return 200 and Hello World', async () => {
        const response = await request(app).get('/');
        expect(response.status).toBe(200);
        expect(response.text).toBe('Hello World');
    });

    // Test for POST /api/chat
    test('POST /api/chat should return 200 and a response', async () => {
        const response = await request(app)
            .post('/api/chat')
            .send({ message: 'Test message' });
        expect(response.status).toBe(200);
        expect(response.body).toHaveProperty('response');
    });

    // Add more tests for other routes...
});