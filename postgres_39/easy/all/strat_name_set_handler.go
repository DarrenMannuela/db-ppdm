package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratNameSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_name_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_name_set

	for rows.Next() {
		var strat_name_set dto.Strat_name_set
		if err := rows.Scan(&strat_name_set.Strat_name_set_id, &strat_name_set.Active_ind, &strat_name_set.Area_id, &strat_name_set.Area_type, &strat_name_set.Business_associate_id, &strat_name_set.Certified_ind, &strat_name_set.Effective_date, &strat_name_set.Expiry_date, &strat_name_set.Ppdm_guid, &strat_name_set.Preferred_ind, &strat_name_set.Remark, &strat_name_set.Source, &strat_name_set.Strat_name_set_name, &strat_name_set.Strat_name_set_type, &strat_name_set.Row_changed_by, &strat_name_set.Row_changed_date, &strat_name_set.Row_created_by, &strat_name_set.Row_created_date, &strat_name_set.Row_effective_date, &strat_name_set.Row_expiry_date, &strat_name_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_name_set to result
		result = append(result, strat_name_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratNameSet(c *fiber.Ctx) error {
	var strat_name_set dto.Strat_name_set

	setDefaults(&strat_name_set)

	if err := json.Unmarshal(c.Body(), &strat_name_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_name_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	strat_name_set.Row_created_date = formatDate(strat_name_set.Row_created_date)
	strat_name_set.Row_changed_date = formatDate(strat_name_set.Row_changed_date)
	strat_name_set.Effective_date = formatDateString(strat_name_set.Effective_date)
	strat_name_set.Expiry_date = formatDateString(strat_name_set.Expiry_date)
	strat_name_set.Row_effective_date = formatDateString(strat_name_set.Row_effective_date)
	strat_name_set.Row_expiry_date = formatDateString(strat_name_set.Row_expiry_date)






	rows, err := stmt.Exec(strat_name_set.Strat_name_set_id, strat_name_set.Active_ind, strat_name_set.Area_id, strat_name_set.Area_type, strat_name_set.Business_associate_id, strat_name_set.Certified_ind, strat_name_set.Effective_date, strat_name_set.Expiry_date, strat_name_set.Ppdm_guid, strat_name_set.Preferred_ind, strat_name_set.Remark, strat_name_set.Source, strat_name_set.Strat_name_set_name, strat_name_set.Strat_name_set_type, strat_name_set.Row_changed_by, strat_name_set.Row_changed_date, strat_name_set.Row_created_by, strat_name_set.Row_created_date, strat_name_set.Row_effective_date, strat_name_set.Row_expiry_date, strat_name_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratNameSet(c *fiber.Ctx) error {
	var strat_name_set dto.Strat_name_set

	setDefaults(&strat_name_set)

	if err := json.Unmarshal(c.Body(), &strat_name_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_name_set.Strat_name_set_id = id

    if strat_name_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_name_set set active_ind = :1, area_id = :2, area_type = :3, business_associate_id = :4, certified_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, preferred_ind = :9, remark = :10, source = :11, strat_name_set_name = :12, strat_name_set_type = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where strat_name_set_id = :21")
	if err != nil {
		return err
	}

	strat_name_set.Row_changed_date = formatDate(strat_name_set.Row_changed_date)
	strat_name_set.Effective_date = formatDateString(strat_name_set.Effective_date)
	strat_name_set.Expiry_date = formatDateString(strat_name_set.Expiry_date)
	strat_name_set.Row_effective_date = formatDateString(strat_name_set.Row_effective_date)
	strat_name_set.Row_expiry_date = formatDateString(strat_name_set.Row_expiry_date)






	rows, err := stmt.Exec(strat_name_set.Active_ind, strat_name_set.Area_id, strat_name_set.Area_type, strat_name_set.Business_associate_id, strat_name_set.Certified_ind, strat_name_set.Effective_date, strat_name_set.Expiry_date, strat_name_set.Ppdm_guid, strat_name_set.Preferred_ind, strat_name_set.Remark, strat_name_set.Source, strat_name_set.Strat_name_set_name, strat_name_set.Strat_name_set_type, strat_name_set.Row_changed_by, strat_name_set.Row_changed_date, strat_name_set.Row_created_by, strat_name_set.Row_effective_date, strat_name_set.Row_expiry_date, strat_name_set.Row_quality, strat_name_set.Strat_name_set_id)
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

func PatchStratNameSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_name_set set "
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
	query += " where strat_name_set_id = :id"

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

func DeleteStratNameSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_name_set dto.Strat_name_set
	strat_name_set.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_name_set where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_name_set.Strat_name_set_id)
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


