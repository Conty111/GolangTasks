/*
Middleware RBAC *

Вам необходимо создать Middleware RoleBasedAuthMiddleware(allowedRoles []string, next http.Handler) http.Handler
для HTTP-сервера, который будет ограничивать доступ к определенным ресурсам в зависимости от роли пользователя.

Middleware должен проверять роль пользователя и разрешать или запрещать доступ к ресурсам.
Роль передается в заголовке X-User-Role

Нужно написать два обработчика для пути "/admin" AdminHandler(w http.ResponseWriter, r *http.Request)
для пути "/user" UserHandler(w http.ResponseWriter, r *http.Request)

Не забудьте про функцию StartRBAC (аналог main):

func StartRBAC() {
    allowedAdminRoles := []string{"admin", "superadmin"}
    allowedUserRoles := []string{"user"}

    // Создание маршрута и применение Middleware для пути "/admin".
    adminHandler := RoleBasedAuthMiddleware(allowedAdminRoles, http.HandlerFunc(AdminHandler))
    http.Handle("/admin", adminHandler)

    // Создание маршрута и применение Middleware для пути "/user".
    userHandler := RoleBasedAuthMiddleware(allowedUserRoles, http.HandlerFunc(UserHandler))
    http.Handle("/user", userHandler)

    // Запуск веб-сервера на порту 8080.
    http.ListenAndServe(":8080", nil)
}

Примечания

Что нужно добавить:
Header: X-User-Role

Ожидаемый результат: Если в запрос будет прокинут X-User-Role - нужно выводить его в таком формате.

Если пользователь - admin:
Admin Resource

Если пользователь - user:
User Resource

Если заголовка нет или роль пользователя не подходит - нужно возвращать http.StatusForbidden
*/

package context

import (
	"context"
	"fmt"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Admin Resource")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Resource")
}

func RoleBasedAuthMiddleware(allowedRoles []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get("X-User-Role")
		for _, role := range allowedRoles {
			if val == role {
				ctx := context.WithValue(r.Context(), "X-User-Role", val)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		w.WriteHeader(http.StatusForbidden)
		return
	})
}

func StartRBAC() {
	allowedAdminRoles := []string{"admin", "superadmin"}
	allowedUserRoles := []string{"user"}

	// Создание маршрута и применение Middleware для пути "/admin".
	adminHandler := RoleBasedAuthMiddleware(allowedAdminRoles, http.HandlerFunc(AdminHandler))
	http.Handle("/admin", adminHandler)

	// Создание маршрута и применение Middleware для пути "/user".
	userHandler := RoleBasedAuthMiddleware(allowedUserRoles, http.HandlerFunc(UserHandler))
	http.Handle("/user", userHandler)

	// Запуск веб-сервера на порту 8080.
	http.ListenAndServe(":8080", nil)
}
