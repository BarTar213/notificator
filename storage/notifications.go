package storage

import "github.com/BarTar213/notificator/models"

func (p *Postgres) GetNotification(notification *models.Notification) error {
	return p.db.Model(notification).WherePK().Select()
}

func (p *Postgres) UpdateNotification(notification *models.Notification) error {
	_, err := p.db.Model(notification).
		Where("userID=?userID").
		Set("read=?read").
		Returning(all).
		Update()

	return err
}

func (p *Postgres) AddNotification(notification *models.Notification) error {
	_, err := p.db.Model(notification).Insert()

	return err
}

func (p *Postgres) DeleteNotification(ID int) error {
	_, err := p.db.Exec("DELETE FROM templates WHERE id=?", ID)

	return err
}

func (p *Postgres) ListNotifications() ([]models.Notification, error) {
	notifications := make([]models.Notification, 0)
	err := p.db.Model(&notifications).Select()
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
