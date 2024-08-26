import { checkAuth, checkAdmin } from './lib/api';

export async function onRequest({ request, redirect }) {
    const url = new URL(request.url);

    if (url.pathname === '/login') {
        return;
    }

    const isAuthenticated = await checkAuth();
    if (!isAuthenticated) {
        return redirect('/login');
    }

    if (url.pathname === '/admin') {
        const isAdmin = await checkAdmin();
        if (!isAdmin) {
            return redirect('/');
        }
    }
}