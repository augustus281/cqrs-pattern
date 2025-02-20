package repository

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/augustus281/cqrs-pattern/global"
	mongomodels "github.com/augustus281/cqrs-pattern/internal/order/models/mongo_models"
	"github.com/augustus281/cqrs-pattern/pkg/constants"
	"github.com/augustus281/cqrs-pattern/pkg/tracing"
)

type mongoRepository struct {
	db *mongo.Client
}

func NewMongoRepository(db *mongo.Client) *mongoRepository {
	return &mongoRepository{
		db: db,
	}
}

func (m *mongoRepository) Insert(ctx context.Context, order *mongomodels.OrderProjection) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.Insert")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	insertOptions := options.InsertOne()
	_, err := m.getOrdersCollection().InsertOne(ctx, order, insertOptions)
	if err != nil {
		tracing.TraceErr(span, err)
		return "", err
	}

	return order.OrderID, nil
}

func (m *mongoRepository) GetByID(ctx context.Context, orderID string) (*mongomodels.OrderProjection, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetByID")
	defer span.Finish()
	span.LogFields(log.String("OrderID", orderID))

	var orderProjection mongomodels.OrderProjection
	err := m.getOrdersCollection().
		FindOne(ctx, bson.M{
			constants.OrderId: orderID,
		}).
		Decode(&orderProjection)
	if err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}

	return &orderProjection, nil
}

func (m *mongoRepository) UpdateOrder(ctx context.Context, order *mongomodels.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateShoppingCart")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	var res mongomodels.OrderProjection
	err := m.getOrdersCollection().
		FindOneAndUpdate(
			ctx,
			bson.M{constants.OrderId: order.OrderID},
			bson.M{"$set": order},
			ops,
		).
		Decode(&res)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	return nil
}

func (m *mongoRepository) UpdateCancel(ctx context.Context, order *mongomodels.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateCanel")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	update := bson.M{
		"$set": bson.M{
			constants.Canceled:     order.Canceled,
			constants.CancelReason: order.CancelReason,
		},
	}

	var res mongomodels.OrderProjection
	err := m.getOrdersCollection().
		FindOneAndUpdate(ctx, bson.M{constants.OrderId: order.OrderID}, update, ops).
		Decode(&res)
	if err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	return nil
}

func (m *mongoRepository) UpdatePayment(ctx context.Context, order *mongomodels.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdatePayment")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	update := bson.M{"$set": bson.M{constants.Payment: order.Payment, constants.Paid: order.Paid}}
	var res mongomodels.OrderProjection
	if err := m.getOrdersCollection().FindOneAndUpdate(ctx, bson.M{constants.OrderId: order.OrderID}, update, ops).Decode(&res); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	global.Logger.Debug(fmt.Sprintf("(UpdatePayment) result OrderID: {%s}", res.OrderID))
	return nil
}

func (m *mongoRepository) Complete(ctx context.Context, order *mongomodels.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.Complete")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	update := bson.M{"$set": bson.M{constants.Completed: order.Completed, constants.DeliveredTime: order.DeliveredTime}}
	var res mongomodels.OrderProjection
	if err := m.getOrdersCollection().FindOneAndUpdate(ctx, bson.M{constants.OrderId: order.OrderID}, update, ops).Decode(&res); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	global.Logger.Debug(fmt.Sprintf("(Complete) result OrderID: {%s}", res.OrderID))
	return nil
}

func (m *mongoRepository) UpdateDeliveryAddress(ctx context.Context, order *mongomodels.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateDeliveryAddress")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	update := bson.M{"$set": bson.M{constants.Completed: order.Completed, constants.DeliveredTime: order.DeliveredTime}}
	var res mongomodels.OrderProjection
	if err := m.getOrdersCollection().FindOneAndUpdate(ctx, bson.M{constants.OrderId: order.OrderID}, update, ops).Decode(&res); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	global.Logger.Debug(fmt.Sprintf("(UpdateDeliveryAddress) result OrderID: {%s}", res.OrderID))
	return nil
}

func (m *mongoRepository) UpdateSubmit(ctx context.Context, order *mongomodels.OrderProjection) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateSubmit")
	defer span.Finish()
	span.LogFields(log.String("OrderID", order.OrderID))

	ops := options.FindOneAndUpdate()
	ops.SetReturnDocument(options.After)
	ops.SetUpsert(false)

	update := bson.M{"$set": bson.M{constants.Submitted: order.Submitted}}
	var res mongomodels.OrderProjection
	if err := m.getOrdersCollection().FindOneAndUpdate(ctx, bson.M{constants.OrderId: order.OrderID}, update, ops).Decode(&res); err != nil {
		tracing.TraceErr(span, err)
		return err
	}

	global.Logger.Debug(fmt.Sprintf("(UpdateSubmit) result OrderID: {%s}", res.OrderID))
	return nil
}

func (m *mongoRepository) getOrdersCollection() *mongo.Collection {
	return m.db.Database(global.Config.MongoDB.Db).Collection(global.Config.MongoDBCollections.Shop)
}
