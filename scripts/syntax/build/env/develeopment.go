//go:build !staging && !product
package env

func GetEnv() string {
	return "development"
}
