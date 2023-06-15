package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossil(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil

	for rows.Next() {
		var fossil dto.Fossil
		if err := rows.Scan(&fossil.Fossil_id, &fossil.Active_ind, &fossil.Climate_type, &fossil.Effective_date, &fossil.Expiry_date, &fossil.Life_habit, &fossil.Lower_ecozone_id, &fossil.Ppdm_guid, &fossil.Preferred_ind, &fossil.Remark, &fossil.Source, &fossil.Taxon_group, &fossil.Taxon_leaf_id, &fossil.Upper_ecozone_id, &fossil.Row_changed_by, &fossil.Row_changed_date, &fossil.Row_created_by, &fossil.Row_created_date, &fossil.Row_effective_date, &fossil.Row_expiry_date, &fossil.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil to result
		result = append(result, fossil)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossil(c *fiber.Ctx) error {
	var fossil dto.Fossil

	setDefaults(&fossil)

	if err := json.Unmarshal(c.Body(), &fossil); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	fossil.Row_created_date = formatDate(fossil.Row_created_date)
	fossil.Row_changed_date = formatDate(fossil.Row_changed_date)
	fossil.Effective_date = formatDateString(fossil.Effective_date)
	fossil.Expiry_date = formatDateString(fossil.Expiry_date)
	fossil.Row_effective_date = formatDateString(fossil.Row_effective_date)
	fossil.Row_expiry_date = formatDateString(fossil.Row_expiry_date)






	rows, err := stmt.Exec(fossil.Fossil_id, fossil.Active_ind, fossil.Climate_type, fossil.Effective_date, fossil.Expiry_date, fossil.Life_habit, fossil.Lower_ecozone_id, fossil.Ppdm_guid, fossil.Preferred_ind, fossil.Remark, fossil.Source, fossil.Taxon_group, fossil.Taxon_leaf_id, fossil.Upper_ecozone_id, fossil.Row_changed_by, fossil.Row_changed_date, fossil.Row_created_by, fossil.Row_created_date, fossil.Row_effective_date, fossil.Row_expiry_date, fossil.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossil(c *fiber.Ctx) error {
	var fossil dto.Fossil

	setDefaults(&fossil)

	if err := json.Unmarshal(c.Body(), &fossil); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil.Fossil_id = id

    if fossil.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil set active_ind = :1, climate_type = :2, effective_date = :3, expiry_date = :4, life_habit = :5, lower_ecozone_id = :6, ppdm_guid = :7, preferred_ind = :8, remark = :9, source = :10, taxon_group = :11, taxon_leaf_id = :12, upper_ecozone_id = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where fossil_id = :21")
	if err != nil {
		return err
	}

	fossil.Row_changed_date = formatDate(fossil.Row_changed_date)
	fossil.Effective_date = formatDateString(fossil.Effective_date)
	fossil.Expiry_date = formatDateString(fossil.Expiry_date)
	fossil.Row_effective_date = formatDateString(fossil.Row_effective_date)
	fossil.Row_expiry_date = formatDateString(fossil.Row_expiry_date)






	rows, err := stmt.Exec(fossil.Active_ind, fossil.Climate_type, fossil.Effective_date, fossil.Expiry_date, fossil.Life_habit, fossil.Lower_ecozone_id, fossil.Ppdm_guid, fossil.Preferred_ind, fossil.Remark, fossil.Source, fossil.Taxon_group, fossil.Taxon_leaf_id, fossil.Upper_ecozone_id, fossil.Row_changed_by, fossil.Row_changed_date, fossil.Row_created_by, fossil.Row_effective_date, fossil.Row_expiry_date, fossil.Row_quality, fossil.Fossil_id)
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

func PatchFossil(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil set "
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
	query += " where fossil_id = :id"

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

func DeleteFossil(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil dto.Fossil
	fossil.Fossil_id = id

	stmt, err := db.Prepare("delete from fossil where fossil_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil.Fossil_id)
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


