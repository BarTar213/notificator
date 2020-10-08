package storage

import "github.com/BarTar213/notificator/models"

func (p *Postgres) AddTemplate(template *models.Template) error {
	_, err := p.db.Model(template).Returning(all).Insert()

	return err
}

func (p *Postgres) GetTemplate(template *models.Template) error {
	return p.db.Model(template).WherePK().Select()
}

func (p *Postgres) UpdateTemplate(template *models.Template) error {
	_, err := p.db.Model(template).
		WherePK().
		Returning(all).
		Update()

	return err
}

func (p *Postgres) DeleteTemplate(ID int) error {
	_, err := p.db.Exec("DELETE FROM templates WHERE id=?", ID)

	return err
}
