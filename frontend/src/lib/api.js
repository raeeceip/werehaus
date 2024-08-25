// src/lib/api.js

const API_BASE_URL = 'http://localhost:3000/api';

// User authentication
export async function login(username, password) {
    const response = await fetch(`${API_BASE_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
    });
    if (!response.ok) throw new Error('Login failed');
    return response.json();
}

export async function logout() {
    const response = await fetch(`${API_BASE_URL}/logout`, { method: 'POST' });
    if (!response.ok) throw new Error('Logout failed');
}

// Item management
export async function fetchItems(page = 1, limit = 10, search = '') {
    const response = await fetch(`${API_BASE_URL}/items?page=${page}&limit=${limit}&search=${search}`);
    if (!response.ok) throw new Error('Failed to fetch items');
    return response.json();
}

export async function createItem(item) {
    const response = await fetch(`${API_BASE_URL}/items`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(item),
    });
    if (!response.ok) throw new Error('Failed to create item');
    return response.json();
}

export async function updateItem(id, item) {
    const response = await fetch(`${API_BASE_URL}/items/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(item),
    });
    if (!response.ok) throw new Error('Failed to update item');
    return response.json();
}

export async function deleteItem(id) {
    const response = await fetch(`${API_BASE_URL}/items/${id}`, { method: 'DELETE' });
    if (!response.ok) throw new Error('Failed to delete item');
}

// Location management
export async function fetchLocations() {
    const response = await fetch(`${API_BASE_URL}/locations`);
    if (!response.ok) throw new Error('Failed to fetch locations');
    return response.json();
}

export async function createLocation(location) {
    const response = await fetch(`${API_BASE_URL}/locations`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(location),
    });
    if (!response.ok) throw new Error('Failed to create location');
    return response.json();
}

export async function updateLocation(id, location) {
    const response = await fetch(`${API_BASE_URL}/locations/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(location),
    });
    if (!response.ok) throw new Error('Failed to update location');
    return response.json();
}

export async function deleteLocation(id) {
    const response = await fetch(`${API_BASE_URL}/locations/${id}`, { method: 'DELETE' });
    if (!response.ok) throw new Error('Failed to delete location');
}

// Issue management
export async function requestIssue(issue) {
    const response = await fetch(`${API_BASE_URL}/issues`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(issue),
    });
    if (!response.ok) throw new Error('Failed to request issue');
    return response.json();
}

export async function fetchPendingIssues() {
    const response = await fetch(`${API_BASE_URL}/issues/pending`);
    if (!response.ok) throw new Error('Failed to fetch pending issues');
    return response.json();
}

export async function approveIssue(id) {
    const response = await fetch(`${API_BASE_URL}/issues/${id}/approve`, { method: 'POST' });
    if (!response.ok) throw new Error('Failed to approve issue');
    return response.json();
}

export async function denyIssue(id) {
    const response = await fetch(`${API_BASE_URL}/issues/${id}/deny`, { method: 'POST' });
    if (!response.ok) throw new Error('Failed to deny issue');
    return response.json();
}

// Reporting
export async function fetchInventoryReport() {
    const response = await fetch(`${API_BASE_URL}/reports/inventory`);
    if (!response.ok) throw new Error('Failed to fetch inventory report');
    return response.json();
}

export async function fetchIssueReport() {
    const response = await fetch(`${API_BASE_URL}/reports/issues`);
    if (!response.ok) throw new Error('Failed to fetch issue report');
    return response.json();
}

export async function fetchItemMovementReport(itemId) {
    const response = await fetch(`${API_BASE_URL}/reports/item-movements/${itemId}`);
    if (!response.ok) throw new Error('Failed to fetch item movement report');
    return response.json();
}