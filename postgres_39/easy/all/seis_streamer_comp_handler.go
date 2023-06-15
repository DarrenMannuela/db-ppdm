package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisStreamerComp(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_streamer_comp")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_streamer_comp

	for rows.Next() {
		var seis_streamer_comp dto.Seis_streamer_comp
		if err := rows.Scan(&seis_streamer_comp.Streamer_id, &seis_streamer_comp.Component_type, &seis_streamer_comp.Active_ind, &seis_streamer_comp.Component_count, &seis_streamer_comp.Effective_date, &seis_streamer_comp.Expiry_date, &seis_streamer_comp.Ppdm_guid, &seis_streamer_comp.Remark, &seis_streamer_comp.Source, &seis_streamer_comp.Row_changed_by, &seis_streamer_comp.Row_changed_date, &seis_streamer_comp.Row_created_by, &seis_streamer_comp.Row_created_date, &seis_streamer_comp.Row_effective_date, &seis_streamer_comp.Row_expiry_date, &seis_streamer_comp.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_streamer_comp to result
		result = append(result, seis_streamer_comp)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisStreamerComp(c *fiber.Ctx) error {
	var seis_streamer_comp dto.Seis_streamer_comp

	setDefaults(&seis_streamer_comp)

	if err := json.Unmarshal(c.Body(), &seis_streamer_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_streamer_comp values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	seis_streamer_comp.Row_created_date = formatDate(seis_streamer_comp.Row_created_date)
	seis_streamer_comp.Row_changed_date = formatDate(seis_streamer_comp.Row_changed_date)
	seis_streamer_comp.Effective_date = formatDateString(seis_streamer_comp.Effective_date)
	seis_streamer_comp.Expiry_date = formatDateString(seis_streamer_comp.Expiry_date)
	seis_streamer_comp.Row_effective_date = formatDateString(seis_streamer_comp.Row_effective_date)
	seis_streamer_comp.Row_expiry_date = formatDateString(seis_streamer_comp.Row_expiry_date)






	rows, err := stmt.Exec(seis_streamer_comp.Streamer_id, seis_streamer_comp.Component_type, seis_streamer_comp.Active_ind, seis_streamer_comp.Component_count, seis_streamer_comp.Effective_date, seis_streamer_comp.Expiry_date, seis_streamer_comp.Ppdm_guid, seis_streamer_comp.Remark, seis_streamer_comp.Source, seis_streamer_comp.Row_changed_by, seis_streamer_comp.Row_changed_date, seis_streamer_comp.Row_created_by, seis_streamer_comp.Row_created_date, seis_streamer_comp.Row_effective_date, seis_streamer_comp.Row_expiry_date, seis_streamer_comp.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisStreamerComp(c *fiber.Ctx) error {
	var seis_streamer_comp dto.Seis_streamer_comp

	setDefaults(&seis_streamer_comp)

	if err := json.Unmarshal(c.Body(), &seis_streamer_comp); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_streamer_comp.Streamer_id = id

    if seis_streamer_comp.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_streamer_comp set component_type = :1, active_ind = :2, component_count = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where streamer_id = :16")
	if err != nil {
		return err
	}

	seis_streamer_comp.Row_changed_date = formatDate(seis_streamer_comp.Row_changed_date)
	seis_streamer_comp.Effective_date = formatDateString(seis_streamer_comp.Effective_date)
	seis_streamer_comp.Expiry_date = formatDateString(seis_streamer_comp.Expiry_date)
	seis_streamer_comp.Row_effective_date = formatDateString(seis_streamer_comp.Row_effective_date)
	seis_streamer_comp.Row_expiry_date = formatDateString(seis_streamer_comp.Row_expiry_date)






	rows, err := stmt.Exec(seis_streamer_comp.Component_type, seis_streamer_comp.Active_ind, seis_streamer_comp.Component_count, seis_streamer_comp.Effective_date, seis_streamer_comp.Expiry_date, seis_streamer_comp.Ppdm_guid, seis_streamer_comp.Remark, seis_streamer_comp.Source, seis_streamer_comp.Row_changed_by, seis_streamer_comp.Row_changed_date, seis_streamer_comp.Row_created_by, seis_streamer_comp.Row_effective_date, seis_streamer_comp.Row_expiry_date, seis_streamer_comp.Row_quality, seis_streamer_comp.Streamer_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchSeisStreamerComp(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_streamer_comp set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where streamer_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteSeisStreamerComp(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_streamer_comp dto.Seis_streamer_comp
	seis_streamer_comp.Streamer_id = id

	stmt, err := db.Prepare("delete from seis_streamer_comp where streamer_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_streamer_comp.Streamer_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}


