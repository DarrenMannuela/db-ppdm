package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlCoalRankScheme(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_coal_rank_scheme")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_coal_rank_scheme

	for rows.Next() {
		var anl_coal_rank_scheme dto.Anl_coal_rank_scheme
		if err := rows.Scan(&anl_coal_rank_scheme.Coal_rank_scheme_id, &anl_coal_rank_scheme.Active_ind, &anl_coal_rank_scheme.Coal_rank_scheme_type, &anl_coal_rank_scheme.Effective_date, &anl_coal_rank_scheme.Expiry_date, &anl_coal_rank_scheme.Ppdm_guid, &anl_coal_rank_scheme.Preferred_ind, &anl_coal_rank_scheme.Remark, &anl_coal_rank_scheme.Scheme_name, &anl_coal_rank_scheme.Scheme_owner_ba_id, &anl_coal_rank_scheme.Source, &anl_coal_rank_scheme.Source_document_id, &anl_coal_rank_scheme.Row_changed_by, &anl_coal_rank_scheme.Row_changed_date, &anl_coal_rank_scheme.Row_created_by, &anl_coal_rank_scheme.Row_created_date, &anl_coal_rank_scheme.Row_effective_date, &anl_coal_rank_scheme.Row_expiry_date, &anl_coal_rank_scheme.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_coal_rank_scheme to result
		result = append(result, anl_coal_rank_scheme)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlCoalRankScheme(c *fiber.Ctx) error {
	var anl_coal_rank_scheme dto.Anl_coal_rank_scheme

	setDefaults(&anl_coal_rank_scheme)

	if err := json.Unmarshal(c.Body(), &anl_coal_rank_scheme); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_coal_rank_scheme values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	anl_coal_rank_scheme.Row_created_date = formatDate(anl_coal_rank_scheme.Row_created_date)
	anl_coal_rank_scheme.Row_changed_date = formatDate(anl_coal_rank_scheme.Row_changed_date)
	anl_coal_rank_scheme.Effective_date = formatDateString(anl_coal_rank_scheme.Effective_date)
	anl_coal_rank_scheme.Expiry_date = formatDateString(anl_coal_rank_scheme.Expiry_date)
	anl_coal_rank_scheme.Row_effective_date = formatDateString(anl_coal_rank_scheme.Row_effective_date)
	anl_coal_rank_scheme.Row_expiry_date = formatDateString(anl_coal_rank_scheme.Row_expiry_date)






	rows, err := stmt.Exec(anl_coal_rank_scheme.Coal_rank_scheme_id, anl_coal_rank_scheme.Active_ind, anl_coal_rank_scheme.Coal_rank_scheme_type, anl_coal_rank_scheme.Effective_date, anl_coal_rank_scheme.Expiry_date, anl_coal_rank_scheme.Ppdm_guid, anl_coal_rank_scheme.Preferred_ind, anl_coal_rank_scheme.Remark, anl_coal_rank_scheme.Scheme_name, anl_coal_rank_scheme.Scheme_owner_ba_id, anl_coal_rank_scheme.Source, anl_coal_rank_scheme.Source_document_id, anl_coal_rank_scheme.Row_changed_by, anl_coal_rank_scheme.Row_changed_date, anl_coal_rank_scheme.Row_created_by, anl_coal_rank_scheme.Row_created_date, anl_coal_rank_scheme.Row_effective_date, anl_coal_rank_scheme.Row_expiry_date, anl_coal_rank_scheme.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlCoalRankScheme(c *fiber.Ctx) error {
	var anl_coal_rank_scheme dto.Anl_coal_rank_scheme

	setDefaults(&anl_coal_rank_scheme)

	if err := json.Unmarshal(c.Body(), &anl_coal_rank_scheme); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_coal_rank_scheme.Coal_rank_scheme_id = id

    if anl_coal_rank_scheme.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_coal_rank_scheme set active_ind = :1, coal_rank_scheme_type = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, preferred_ind = :6, remark = :7, scheme_name = :8, scheme_owner_ba_id = :9, source = :10, source_document_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where coal_rank_scheme_id = :19")
	if err != nil {
		return err
	}

	anl_coal_rank_scheme.Row_changed_date = formatDate(anl_coal_rank_scheme.Row_changed_date)
	anl_coal_rank_scheme.Effective_date = formatDateString(anl_coal_rank_scheme.Effective_date)
	anl_coal_rank_scheme.Expiry_date = formatDateString(anl_coal_rank_scheme.Expiry_date)
	anl_coal_rank_scheme.Row_effective_date = formatDateString(anl_coal_rank_scheme.Row_effective_date)
	anl_coal_rank_scheme.Row_expiry_date = formatDateString(anl_coal_rank_scheme.Row_expiry_date)






	rows, err := stmt.Exec(anl_coal_rank_scheme.Active_ind, anl_coal_rank_scheme.Coal_rank_scheme_type, anl_coal_rank_scheme.Effective_date, anl_coal_rank_scheme.Expiry_date, anl_coal_rank_scheme.Ppdm_guid, anl_coal_rank_scheme.Preferred_ind, anl_coal_rank_scheme.Remark, anl_coal_rank_scheme.Scheme_name, anl_coal_rank_scheme.Scheme_owner_ba_id, anl_coal_rank_scheme.Source, anl_coal_rank_scheme.Source_document_id, anl_coal_rank_scheme.Row_changed_by, anl_coal_rank_scheme.Row_changed_date, anl_coal_rank_scheme.Row_created_by, anl_coal_rank_scheme.Row_effective_date, anl_coal_rank_scheme.Row_expiry_date, anl_coal_rank_scheme.Row_quality, anl_coal_rank_scheme.Coal_rank_scheme_id)
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

func PatchAnlCoalRankScheme(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_coal_rank_scheme set "
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
	query += " where coal_rank_scheme_id = :id"

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

func DeleteAnlCoalRankScheme(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_coal_rank_scheme dto.Anl_coal_rank_scheme
	anl_coal_rank_scheme.Coal_rank_scheme_id = id

	stmt, err := db.Prepare("delete from anl_coal_rank_scheme where coal_rank_scheme_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_coal_rank_scheme.Coal_rank_scheme_id)
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


