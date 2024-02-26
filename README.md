# Go-Backend Coffee Shop

Welcome to the Coffee Shop Backend Web Project! This repository contains the back-end source code for the Online Coffee Shop web application. With Golang structure.

Find a wide range of selected coffee, from coffee beans, ground coffee, to quality coffee equipment here. Order your favorite coffee easily and enjoy the best coffee experience at home.


## Features

- Explore a wide range of coffee products from various local and international roasters.
- Read complete product descriptions to help you choose the right coffee.
- Order coffee easily and safely through a trusted online payment system.
- Get exciting promos and special offers.
- Easily monitor your order status.

Built of

![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Nodemon](https://img.shields.io/badge/NODEMON-%23323330.svg?style=for-the-badge&logo=nodemon&logoColor=%BBDEAD)
![Go](https://img.shields.io/badge/Go-blue?style=for-the-badge&logo=go&logoColor=white)
![gin-gonic](https://img.shields.io/badge/gin-1DA1F2?style=for-the-badge&logo=gin&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![NPM](https://img.shields.io/badge/NPM-%23CB3837.svg?style=for-the-badge&logo=npm&logoColor=white)
![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white)
![Google Chrome](https://img.shields.io/badge/Google%20Chrome-4285F4?style=for-the-badge&logo=GoogleChrome&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)


## 📌 Getting Started

To run the project locally, follow these simple steps:

1. Clone this repository
```sh
  git clone https://github.com/DzulfiqarSiraj/fwg17-go-backend.git
  cd fwg17-go-backend
```

2. Open in VSCode
```sh
  code .
```

3. install all the dependencies
```sh
  go mod tidy
```

4. run the project
```sh
  go run .
```

5. run the project using nodemon
```sh
  npm install nodemone
```
```sh
  nodemon --exec go run main.go --signal SIGTERM


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`POSTGRES_URL`

`APP_SECRET`


## API Reference

#### Login

```http
  POST /login
```
#### Register

```http
  POST /register
```
#### Forgot Password

```http
  POST /forgot-password
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `users` | `GET` | Get a list of users data |
| `users/:id` | `GET` | Get a detailed users data |
| `users` | `POST` | Insert a users data |
| `users/:id` | `PATCH` | Update a users data |
| `users/:id` | `DELETE` | Delete a users data |
| `products` | `GET` | Get a list of products data |
| `products/:id` | `GET` | Get a detailed products data |
| `products` | `POST` | Insert a products data |
| `products/:id` | `PATCH` | Update a products data |
| `products/:id` | `DELETE` | Delete a products data |
| `promo` | `GET` | Get a list of promo data |
| `promo/:id` | `GET` | Get a detailed promo data |
| `promo` | `POST` | Insert a promo data |
| `promo/:id` | `PATCH` | Update a promo data |
| `promo/:id` | `DELETE` | Delete a promo data |
| `product-size` | `GET` | Get a list of product size data |
| `product-size/:id` | `GET` | Get a detailed product size data |
| `product-size/:id` | `PATCH` | Update a product size data |
| `product-variant` | `GET` | Get a list of product variant data |
| `product-variant/:id` | `GET` | Get a detailed product variant data |
| `product-variant` | `POST` | Insert a product variant data |
| `product-variant/:id` | `PATCH` | Update a product variant data |
| `product-variant/:id` | `DELETE` | Delete a product variant data |
| `categories` | `GET` | Get a list of categories data |
| `categories/:id` | `GET` | Get a detailed categories data |
| `categories` | `POST` | Insert a categories data |
| `categories/:id` | `PATCH` | Update a categories data |
| `categories/:id` | `DELETE` | Delete a categories data |
| `tags` | `GET` | Get a list of tags data |
| `tags/:id` | `GET` | Get a detailed tags data |
| `tags` | `POST` | Insert a tags data |
| `tags/:id` | `PATCH` | Update a tags data |
| `tags/:id` | `DELETE` | Delete a tags data |
| `product-tags` | `GET` | Get a list of product tags data |
| `product-tags/:id` | `GET` | Get a detailed product tags data |
| `product-tags` | `POST` | Insert a product tags data |
| `product-tags/:id` | `PATCH` | Update a product tags data |
| `product-tags/:id` | `DELETE` | Delete a product tags data |
| `product-categories` | `GET` | Get a list of product categories data |
| `product-categories/:id` | `GET` | Get a detailed product categories data |
| `product-categories` | `POST` | Insert a product categories data |
| `product-categories/:id` | `PATCH` | Update a product categories data |
| `product-categories/:id` | `DELETE` | Delete a product categories data |
| `product-ratings` | `GET` | Get a list of product ratings data |
| `product-ratings/:id` | `GET` | Get a detailed product ratings data |
| `product-ratings` | `POST` | Insert a product ratings data |
| `product-ratings/:id` | `PATCH` | Update a product ratings data |
| `product-ratings/:id` | `DELETE` | Delete a product ratings data |
| `cart` | `GET` | Get a list of cart data |
| `cart/:id` | `GET` | Get a detailed cart data |
| `cart` | `POST` | Insert a cart data |
| `cart/:id` | `PATCH` | Update a cart data |
| `cart/:id` | `DELETE` | Delete a cart data |
| `orders` | `GET` | Get a list of orders data |
| `orders/:id` | `GET` | Get a detailed orders data |
| `orders` | `POST` | Insert a orders data |
| `orders/:id` | `PATCH` | Update a orders data |
| `orders/:id` | `DELETE` | Delete a orders data |
| `order-details` | `GET` | Get a list of order details data |
| `order-details/:id` | `GET` | Get a detailed order details data |
| `order-details` | `POST` | Insert a order details data |
| `order-details/:id` | `PATCH` | Update a order details data |
| `order-details/:id` | `DELETE` | Delete a order details data |

## Technologies Used

**Gin-Gonic:** a fast and minimalist web framework, this effort is dedicated to creating fast as well as resilient and scalable server-side applications.

**Go:** Go serves as the foundation for this project, utilizing a non-blocking and event-driven architecture to ensure scalable and high-performance server-side application development.

## ✍️ Coffee Shop - Frontend Repository
https://github.com/DzulfiqarSiraj/fwg17-beginner-frontend

## Technologies Used

The project structure is organized as follows: 
- src/: contains the source code of the project.
- asset/: image and icon media.
- components/: Reusable component used throughout the project.
- pages/: Individual pages of the application.


## Contributing

Contributions are always welcome!

## Authors

- [@DzulfiqarSiraj](https://github.com/DzulfiqarSiraj)

## Feedback

If you have any feedback, please reach out to us at shalahuddindzulfiqar@gmail.com
