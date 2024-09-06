import { checkAuth, checkAdmin } from './lib/api';

export async function onRequest({ request, redirect }) {
    const url = new URL(request.url);

    // Allow access to login page without authentication
    if (url.pathname === '/login') {
        return;
    }

    // Check if user is authenticated
    const isAuthenticated = await checkAuth();
    if (!isAuthenticated) {
        return redirect('/login');
    }

    // Check if user is admin for admin-only pages
    if (url.pathname === '/admin') {
        const isAdmin = await checkAdmin();
        if (!isAdmin) {
            return redirect('/');
        }
    }
}