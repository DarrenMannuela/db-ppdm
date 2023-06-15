package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoAbundScheme(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_abund_scheme")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_abund_scheme

	for rows.Next() {
		var paleo_abund_scheme dto.Paleo_abund_scheme
		if err := rows.Scan(&paleo_abund_scheme.Scheme_id, &paleo_abund_scheme.Abund_qualifier_id, &paleo_abund_scheme.Active_ind, &paleo_abund_scheme.Effective_date, &paleo_abund_scheme.Expiry_date, &paleo_abund_scheme.Max_count, &paleo_abund_scheme.Min_count, &paleo_abund_scheme.Ppdm_guid, &paleo_abund_scheme.Remark, &paleo_abund_scheme.Scheme_owner_ba_id, &paleo_abund_scheme.Source, &paleo_abund_scheme.Row_changed_by, &paleo_abund_scheme.Row_changed_date, &paleo_abund_scheme.Row_created_by, &paleo_abund_scheme.Row_created_date, &paleo_abund_scheme.Row_effective_date, &paleo_abund_scheme.Row_expiry_date, &paleo_abund_scheme.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_abund_scheme to result
		result = append(result, paleo_abund_scheme)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoAbundScheme(c *fiber.Ctx) error {
	var paleo_abund_scheme dto.Paleo_abund_scheme

	setDefaults(&paleo_abund_scheme)

	if err := json.Unmarshal(c.Body(), &paleo_abund_scheme); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_abund_scheme values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	paleo_abund_scheme.Row_created_date = formatDate(paleo_abund_scheme.Row_created_date)
	paleo_abund_scheme.Row_changed_date = formatDate(paleo_abund_scheme.Row_changed_date)
	paleo_abund_scheme.Effective_date = formatDateString(paleo_abund_scheme.Effective_date)
	paleo_abund_scheme.Expiry_date = formatDateString(paleo_abund_scheme.Expiry_date)
	paleo_abund_scheme.Row_effective_date = formatDateString(paleo_abund_scheme.Row_effective_date)
	paleo_abund_scheme.Row_expiry_date = formatDateString(paleo_abund_scheme.Row_expiry_date)






	rows, err := stmt.Exec(paleo_abund_scheme.Scheme_id, paleo_abund_scheme.Abund_qualifier_id, paleo_abund_scheme.Active_ind, paleo_abund_scheme.Effective_date, paleo_abund_scheme.Expiry_date, paleo_abund_scheme.Max_count, paleo_abund_scheme.Min_count, paleo_abund_scheme.Ppdm_guid, paleo_abund_scheme.Remark, paleo_abund_scheme.Scheme_owner_ba_id, paleo_abund_scheme.Source, paleo_abund_scheme.Row_changed_by, paleo_abund_scheme.Row_changed_date, paleo_abund_scheme.Row_created_by, paleo_abund_scheme.Row_created_date, paleo_abund_scheme.Row_effective_date, paleo_abund_scheme.Row_expiry_date, paleo_abund_scheme.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoAbundScheme(c *fiber.Ctx) error {
	var paleo_abund_scheme dto.Paleo_abund_scheme

	setDefaults(&paleo_abund_scheme)

	if err := json.Unmarshal(c.Body(), &paleo_abund_scheme); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_abund_scheme.Scheme_id = id

    if paleo_abund_scheme.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_abund_scheme set abund_qualifier_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, max_count = :5, min_count = :6, ppdm_guid = :7, remark = :8, scheme_owner_ba_id = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where scheme_id = :18")
	if err != nil {
		return err
	}

	paleo_abund_scheme.Row_changed_date = formatDate(paleo_abund_scheme.Row_changed_date)
	paleo_abund_scheme.Effective_date = formatDateString(paleo_abund_scheme.Effective_date)
	paleo_abund_scheme.Expiry_date = formatDateString(paleo_abund_scheme.Expiry_date)
	paleo_abund_scheme.Row_effective_date = formatDateString(paleo_abund_scheme.Row_effective_date)
	paleo_abund_scheme.Row_expiry_date = formatDateString(paleo_abund_scheme.Row_expiry_date)






	rows, err := stmt.Exec(paleo_abund_scheme.Abund_qualifier_id, paleo_abund_scheme.Active_ind, paleo_abund_scheme.Effective_date, paleo_abund_scheme.Expiry_date, paleo_abund_scheme.Max_count, paleo_abund_scheme.Min_count, paleo_abund_scheme.Ppdm_guid, paleo_abund_scheme.Remark, paleo_abund_scheme.Scheme_owner_ba_id, paleo_abund_scheme.Source, paleo_abund_scheme.Row_changed_by, paleo_abund_scheme.Row_changed_date, paleo_abund_scheme.Row_created_by, paleo_abund_scheme.Row_effective_date, paleo_abund_scheme.Row_expiry_date, paleo_abund_scheme.Row_quality, paleo_abund_scheme.Scheme_id)
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

func PatchPaleoAbundScheme(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_abund_scheme set "
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
	query += " where scheme_id = :id"

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

func DeletePaleoAbundScheme(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_abund_scheme dto.Paleo_abund_scheme
	paleo_abund_scheme.Scheme_id = id

	stmt, err := db.Prepare("delete from paleo_abund_scheme where scheme_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_abund_scheme.Scheme_id)
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


