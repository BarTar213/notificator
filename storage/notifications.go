package storage

import "github.com/BarTar213/notificator/models"

func (p *Postgres) GetNotification(notification *models.Notification) error {
	return p.db.Model(notification).
		WherePK().
		Where("user_id=?user_id").
		Select()
}

func (p *Postgres) ReadNotification(id, userID int, read bool) error {
	_, err := p.db.ExecOne("UPDATE notifications SET read=? WHERE id=? AND user_id=?", read, id, userID)

	return err
}

func (p *Postgres) AddNotification(notification *models.Notification) error {
	_, err := p.db.Model(notification).Insert()

	return err
}

func (p *Postgres) BatchAddNotifications(notifications []*models.Notification) error {
	_, err := p.db.Model(&notifications).Insert()

	return err
}

func (p *Postgres) DeleteNotification(notificationID, userID int) error {
	_, err := p.db.Exec("DELETE FROM templates WHERE id=? AND user_id=?", notificationID, userID)

	return err
}

func (p *Postgres) ListNotifications(userID int) ([]models.Notification, error) {
	notifications := make([]models.Notification, 0)
	err := p.db.Model(&notifications).
		Where("user_id=?", userID).
		Select()
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
