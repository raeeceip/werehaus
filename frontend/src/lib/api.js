// src/lib/api.js

const API_BASE_URL = 'http://localhost:3000/api';

// User authentication
export async function login(username, password) {
    const response = await fetch(`${API_BASE_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
        credentials: 'include', // This is important for including cookies
    });
    if (!response.ok) throw new Error('Login failed');
    return response.json();
}

export async function logout() {
    const response = await fetch(`${API_BASE_URL}/logout`, {
        method: 'POST',
        credentials: 'include',
    });
    if (!response.ok) throw new Error('Logout failed');
}

// Helper function to make authenticated requests
async function authenticatedFetch(url, options = {}) {
    const response = await fetch(url, {
        ...options,
        credentials: 'include',
    });
    if (response.status === 401) {
        // Redirect to login page if unauthorized
        window.location.href = '/login';
        throw new Error('Unauthorized');
    }
    if (!response.ok) throw new Error('Request failed');
    return response;
}

// Item management
export async function fetchItems(page = 1, limit = 10, search = '') {
    const response = await authenticatedFetch(`${API_BASE_URL}/items?page=${page}&limit=${limit}&search=${search}`);
    return response.json();
}

export async function createItem(item) {
    const response = await authenticatedFetch(`${API_BASE_URL}/items`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(item),
    });
    return response.json();
}

export async function updateItem(id, item) {
    const response = await authenticatedFetch(`${API_BASE_URL}/items/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(item),
    });
    return response.json();
}

export async function deleteItem(id) {
    await authenticatedFetch(`${API_BASE_URL}/items/${id}`, { method: 'DELETE' });
}

// Location management
export async function fetchLocations() {
    const response = await authenticatedFetch(`${API_BASE_URL}/locations`);
    return response.json();
}

export async function createLocation(location) {
    const response = await authenticatedFetch(`${API_BASE_URL}/locations`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(location),
    });
    return response.json();
}

export async function updateLocation(id, location) {
    const response = await authenticatedFetch(`${API_BASE_URL}/locations/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(location),
    });
    return response.json();
}

export async function deleteLocation(id) {
    await authenticatedFetch(`${API_BASE_URL}/locations/${id}`, { method: 'DELETE' });
}

// Issue management
export async function requestIssue(issue) {
    const response = await authenticatedFetch(`${API_BASE_URL}/issues`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(issue),
    });
    return response.json();
}

export async function fetchPendingIssues() {
    const response = await authenticatedFetch(`${API_BASE_URL}/issues/pending`);
    return response.json();
}

export async function approveIssue(id) {
    const response = await authenticatedFetch(`${API_BASE_URL}/issues/${id}/approve`, { method: 'POST' });
    return response.json();
}

export async function denyIssue(id) {
    const response = await authenticatedFetch(`${API_BASE_URL}/issues/${id}/deny`, { method: 'POST' });
    return response.json();
}

// Reporting
export async function fetchInventoryReport() {
    const response = await authenticatedFetch(`${API_BASE_URL}/reports/inventory`);
    return response.json();
}

export async function fetchIssueReport() {
    const response = await authenticatedFetch(`${API_BASE_URL}/reports/issues`);
    return response.json();
}

export async function fetchItemMovementReport(itemId) {
    const response = await authenticatedFetch(`${API_BASE_URL}/reports/item-movements/${itemId}`);
    return response.json();
}

// Admin-only functions
export async function createUser(user) {
    const response = await authenticatedFetch(`${API_BASE_URL}/admin/users`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user),
    });
    return response.json();
}

export async function fetchUsers() {
    const response = await authenticatedFetch(`${API_BASE_URL}/admin/users`);
    return response.json();
}

// Function to check if the user is authenticated
export async function checkAuth() {
    try {
        await authenticatedFetch(`${API_BASE_URL}/items`);
        return true;
    } catch (error) {
        return false;
    }
}

// Function to check if the user is an admin
export async function checkAdmin() {
    try {
        await authenticatedFetch(`${API_BASE_URL}/admin/users`);
        return true;
    } catch (error) {
        return false;
    }
}