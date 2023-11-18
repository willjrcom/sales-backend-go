/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"flag"

	"github.com/spf13/cobra"
	"github.com/willjrcom/sales-backend-go/bootstrap/database"
	"github.com/willjrcom/sales-backend-go/bootstrap/server"
	handlerimpl "github.com/willjrcom/sales-backend-go/internal/infra/handler"
	addressrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/address"
	categoryrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/category_product"
	clientrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/client"
	contactrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/contact"
	employeerepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/employee"
	groupitemrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/group_item"
	itemrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/item"
	orderrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/order"
	productrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/product"
	quantityrepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/quantity_category"
	sizerepositorybun "github.com/willjrcom/sales-backend-go/internal/infra/repository/postgres/size_category"
	categoryproductusecases "github.com/willjrcom/sales-backend-go/internal/usecases/category_product"
	clientusecases "github.com/willjrcom/sales-backend-go/internal/usecases/client"
	contactusecases "github.com/willjrcom/sales-backend-go/internal/usecases/contact_person"
	deliveryorderusecases "github.com/willjrcom/sales-backend-go/internal/usecases/delivery_order"
	groupitemusecases "github.com/willjrcom/sales-backend-go/internal/usecases/group_item"
	itemusecases "github.com/willjrcom/sales-backend-go/internal/usecases/item"
	orderusecases "github.com/willjrcom/sales-backend-go/internal/usecases/order"
	productusecases "github.com/willjrcom/sales-backend-go/internal/usecases/product"
	quantityusecases "github.com/willjrcom/sales-backend-go/internal/usecases/quantity_category"
	sizeusecases "github.com/willjrcom/sales-backend-go/internal/usecases/size_category"
)

// httpserverCmd represents the httpserver command
var HttpserverCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println("httpserver called")
		port, _ := cmd.Flags().GetString("port")

		flag.Parse()
		ctx := context.Background()
		server := server.NewServerChi()

		// Load database
		db, err := database.NewPostgreSQLConnection(ctx)

		if err != nil {
			panic(err)
		}

		// Load repositories
		productRepo := productrepositorybun.NewProductRepositoryBun(db)
		categoryRepo := categoryrepositorybun.NewCategoryProductRepositoryBun(db)
		sizeRepo := sizerepositorybun.NewSizeCategoryRepositoryBun(db)
		quantityRepo := quantityrepositorybun.NewQuantityCategoryRepositoryBun(db)

		clientRepo := clientrepositorybun.NewClientRepositoryBun(db)
		contactRepo := contactrepositorybun.NewContactRepositoryBun(db)
		addressRepo := addressrepositorybun.NewAddressRepositoryBun(db)

		orderRepo := orderrepositorybun.NewOrderRepositoryBun(db)
		deliveryOrderRepo := orderrepositorybun.NewDeliveryOrderRepositoryBun(db)
		itemRepo := itemrepositorybun.NewItemRepositoryBun(db)
		groupItemRepo := groupitemrepositorybun.NewGroupItemRepositoryBun(db)

		employeeRepo := employeerepositorybun.NewProductRepositoryBun(db)

		// Load services
		productService := productusecases.NewService(productRepo, categoryRepo)
		categoryProductService := categoryproductusecases.NewService(categoryRepo)
		sizeService := sizeusecases.NewService(sizeRepo)
		quantityService := quantityusecases.NewService(quantityRepo)

		clientService := clientusecases.NewService(clientRepo, contactRepo)
		contactService := contactusecases.NewService(contactRepo)

		orderService := orderusecases.NewService(orderRepo)
		deliveryOrderService := deliveryorderusecases.NewService(deliveryOrderRepo, addressRepo, clientRepo, orderRepo, employeeRepo)
		itemService := itemusecases.NewService(itemRepo, groupItemRepo, orderRepo, productRepo)
		groupService := groupitemusecases.NewService(itemRepo, groupItemRepo)

		// Load handlers
		productHandler := handlerimpl.NewHandlerProduct(productService)
		categoryHandler := handlerimpl.NewHandlerCategoryProduct(categoryProductService)
		sizeHandler := handlerimpl.NewHandlerSizeProduct(sizeService)
		quantityHandler := handlerimpl.NewHandlerQuantityProduct(quantityService)

		clientHandler := handlerimpl.NewHandlerClient(clientService)
		contactHandler := handlerimpl.NewHandlerContactPerson(contactService)

		orderHandler := handlerimpl.NewHandlerOrder(orderService)
		deliveryOrderHandler := handlerimpl.NewHandlerDeliveryOrder(deliveryOrderService)
		itemHandler := handlerimpl.NewHandlerItem(itemService)
		groupHandler := handlerimpl.NewHandlerGroupItem(groupService)

		server.AddHandler(productHandler)
		server.AddHandler(categoryHandler)
		server.AddHandler(sizeHandler)
		server.AddHandler(quantityHandler)

		server.AddHandler(clientHandler)
		server.AddHandler(contactHandler)

		server.AddHandler(orderHandler)
		server.AddHandler(deliveryOrderHandler)
		server.AddHandler(itemHandler)
		server.AddHandler(groupHandler)

		if err := server.StartServer(port); err != nil {
			panic(err)
		}
	},
}
