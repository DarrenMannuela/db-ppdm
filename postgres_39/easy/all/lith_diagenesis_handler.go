package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithDiagenesis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_diagenesis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_diagenesis

	for rows.Next() {
		var lith_diagenesis dto.Lith_diagenesis
		if err := rows.Scan(&lith_diagenesis.Lithology_log_id, &lith_diagenesis.Lith_log_source, &lith_diagenesis.Depth_obs_no, &lith_diagenesis.Rock_type, &lith_diagenesis.Rock_type_obs_no, &lith_diagenesis.Diagenesis_type, &lith_diagenesis.Active_ind, &lith_diagenesis.Diagenesis_percent, &lith_diagenesis.Diagenesis_rel_abundance, &lith_diagenesis.Effective_date, &lith_diagenesis.Expiry_date, &lith_diagenesis.Ppdm_guid, &lith_diagenesis.Remark, &lith_diagenesis.Row_changed_by, &lith_diagenesis.Row_changed_date, &lith_diagenesis.Row_created_by, &lith_diagenesis.Row_created_date, &lith_diagenesis.Row_effective_date, &lith_diagenesis.Row_expiry_date, &lith_diagenesis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_diagenesis to result
		result = append(result, lith_diagenesis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithDiagenesis(c *fiber.Ctx) error {
	var lith_diagenesis dto.Lith_diagenesis

	setDefaults(&lith_diagenesis)

	if err := json.Unmarshal(c.Body(), &lith_diagenesis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_diagenesis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	lith_diagenesis.Row_created_date = formatDate(lith_diagenesis.Row_created_date)
	lith_diagenesis.Row_changed_date = formatDate(lith_diagenesis.Row_changed_date)
	lith_diagenesis.Effective_date = formatDateString(lith_diagenesis.Effective_date)
	lith_diagenesis.Expiry_date = formatDateString(lith_diagenesis.Expiry_date)
	lith_diagenesis.Row_effective_date = formatDateString(lith_diagenesis.Row_effective_date)
	lith_diagenesis.Row_expiry_date = formatDateString(lith_diagenesis.Row_expiry_date)






	rows, err := stmt.Exec(lith_diagenesis.Lithology_log_id, lith_diagenesis.Lith_log_source, lith_diagenesis.Depth_obs_no, lith_diagenesis.Rock_type, lith_diagenesis.Rock_type_obs_no, lith_diagenesis.Diagenesis_type, lith_diagenesis.Active_ind, lith_diagenesis.Diagenesis_percent, lith_diagenesis.Diagenesis_rel_abundance, lith_diagenesis.Effective_date, lith_diagenesis.Expiry_date, lith_diagenesis.Ppdm_guid, lith_diagenesis.Remark, lith_diagenesis.Row_changed_by, lith_diagenesis.Row_changed_date, lith_diagenesis.Row_created_by, lith_diagenesis.Row_created_date, lith_diagenesis.Row_effective_date, lith_diagenesis.Row_expiry_date, lith_diagenesis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithDiagenesis(c *fiber.Ctx) error {
	var lith_diagenesis dto.Lith_diagenesis

	setDefaults(&lith_diagenesis)

	if err := json.Unmarshal(c.Body(), &lith_diagenesis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_diagenesis.Lithology_log_id = id

    if lith_diagenesis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_diagenesis set lith_log_source = :1, depth_obs_no = :2, rock_type = :3, rock_type_obs_no = :4, diagenesis_type = :5, active_ind = :6, diagenesis_percent = :7, diagenesis_rel_abundance = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where lithology_log_id = :20")
	if err != nil {
		return err
	}

	lith_diagenesis.Row_changed_date = formatDate(lith_diagenesis.Row_changed_date)
	lith_diagenesis.Effective_date = formatDateString(lith_diagenesis.Effective_date)
	lith_diagenesis.Expiry_date = formatDateString(lith_diagenesis.Expiry_date)
	lith_diagenesis.Row_effective_date = formatDateString(lith_diagenesis.Row_effective_date)
	lith_diagenesis.Row_expiry_date = formatDateString(lith_diagenesis.Row_expiry_date)






	rows, err := stmt.Exec(lith_diagenesis.Lith_log_source, lith_diagenesis.Depth_obs_no, lith_diagenesis.Rock_type, lith_diagenesis.Rock_type_obs_no, lith_diagenesis.Diagenesis_type, lith_diagenesis.Active_ind, lith_diagenesis.Diagenesis_percent, lith_diagenesis.Diagenesis_rel_abundance, lith_diagenesis.Effective_date, lith_diagenesis.Expiry_date, lith_diagenesis.Ppdm_guid, lith_diagenesis.Remark, lith_diagenesis.Row_changed_by, lith_diagenesis.Row_changed_date, lith_diagenesis.Row_created_by, lith_diagenesis.Row_effective_date, lith_diagenesis.Row_expiry_date, lith_diagenesis.Row_quality, lith_diagenesis.Lithology_log_id)
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

func PatchLithDiagenesis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_diagenesis set "
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
	query += " where lithology_log_id = :id"

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

func DeleteLithDiagenesis(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_diagenesis dto.Lith_diagenesis
	lith_diagenesis.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_diagenesis where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_diagenesis.Lithology_log_id)
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


