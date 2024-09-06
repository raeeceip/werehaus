const API_BASE_URL = "http://localhost:3000/api";

export async function checkAuth() {
    // Implement your authentication check logic here
    // For example, you might check for a valid token in localStorage
    const token = localStorage.getItem('authToken');
    if (!token) return false;

    try {
        const response = await fetch(`${API_BASE_URL}/verify-token`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        return response.ok;
    } catch (error) {
        console.error('Auth check failed:', error);
        return false;
    }
}

export async function checkAdmin() {
    // Implement your admin check logic here
    const token = localStorage.getItem('authToken');
    if (!token) return false;

    try {
        const response = await fetch(`${API_BASE_URL}/check-admin`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        return response.ok;
    } catch (error) {
        console.error('Admin check failed:', error);
        return false;
    }
}

export async function loginUser(credentials) {
    const response = await fetch(`${API_BASE_URL}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(credentials),
    });
    const data = await response.json();
    if (data.token) {
        localStorage.setItem('authToken', data.token);
    }
    return data;
}

export async function createUser(userData) {
    const response = await fetch(`${API_BASE_URL}/signup`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
    });
    return response.json();
}

export async function fetchItems() {
    const response = await fetch(`${API_BASE_URL}/items`);
    return response.json();
}

export async function createItem(item) {
    const response = await fetch(`${API_BASE_URL}/items`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(item),
    });
    return response.json();
}

export async function updateItem(id, item) {
    const response = await fetch(`${API_BASE_URL}/items/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(item),
    });
    return response.json();
}

export async function deleteItem(id) {
    const response = await fetch(`${API_BASE_URL}/items/${id}`, {
        method: "DELETE",
    });
    return response.json();
}

export async function fetchLocations() {
    const response = await fetch(`${API_BASE_URL}/locations`);
    return response.json();
}

export async function createLocation(location) {
    const response = await fetch(`${API_BASE_URL}/locations`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(location),
    });
    return response.json();
}

export async function updateLocation(id, location) {
    const response = await fetch(`${API_BASE_URL}/locations/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(location),
    });
    return response.json();
}

export async function deleteLocation(id) {
    const response = await fetch(`${API_BASE_URL}/locations/${id}`, {
        method: "DELETE",
    });
    return response.json();
}

export async function createIssue(issue) {
    const response = await fetch(`${API_BASE_URL}/issues`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(issue),
    });
    return response.json();
}

export async function fetchPendingIssues() {
    const response = await fetch(`${API_BASE_URL}/issues/pending`);
    return response.json();
}

export async function approveIssue(id) {
    const response = await fetch(`${API_BASE_URL}/issues/${id}/approve`, {
        method: "POST",
    });
    return response.json();
}

export async function denyIssue(id) {
    const response = await fetch(`${API_BASE_URL}/issues/${id}/deny`, {
        method: "POST",
    });
    return response.json();
}

export async function fetchInventoryReport() {
    const response = await fetch(`${API_BASE_URL}/reports/inventory`);
    return response.json();
}

export async function fetchIssueReport() {
    const response = await fetch(`${API_BASE_URL}/reports/issues`);
    return response.json();
}

export async function fetchItemMovementReport(itemId) {
    const response = await fetch(`${API_BASE_URL}/reports/item-movements/${itemId}`);
    return response.json();
}