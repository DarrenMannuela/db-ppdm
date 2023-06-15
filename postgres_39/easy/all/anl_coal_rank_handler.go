package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlCoalRank(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_coal_rank")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_coal_rank

	for rows.Next() {
		var anl_coal_rank dto.Anl_coal_rank
		if err := rows.Scan(&anl_coal_rank.Coal_rank_id, &anl_coal_rank.Abbreviation, &anl_coal_rank.Active_ind, &anl_coal_rank.Coal_rank_scheme_id, &anl_coal_rank.Effective_date, &anl_coal_rank.Expiry_date, &anl_coal_rank.Long_name, &anl_coal_rank.Max_value, &anl_coal_rank.Max_value_ouom, &anl_coal_rank.Max_value_uom, &anl_coal_rank.Min_value, &anl_coal_rank.Min_value_ouom, &anl_coal_rank.Min_value_uom, &anl_coal_rank.Ppdm_guid, &anl_coal_rank.Remark, &anl_coal_rank.Short_name, &anl_coal_rank.Source, &anl_coal_rank.Row_changed_by, &anl_coal_rank.Row_changed_date, &anl_coal_rank.Row_created_by, &anl_coal_rank.Row_created_date, &anl_coal_rank.Row_effective_date, &anl_coal_rank.Row_expiry_date, &anl_coal_rank.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_coal_rank to result
		result = append(result, anl_coal_rank)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlCoalRank(c *fiber.Ctx) error {
	var anl_coal_rank dto.Anl_coal_rank

	setDefaults(&anl_coal_rank)

	if err := json.Unmarshal(c.Body(), &anl_coal_rank); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_coal_rank values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	anl_coal_rank.Row_created_date = formatDate(anl_coal_rank.Row_created_date)
	anl_coal_rank.Row_changed_date = formatDate(anl_coal_rank.Row_changed_date)
	anl_coal_rank.Effective_date = formatDateString(anl_coal_rank.Effective_date)
	anl_coal_rank.Expiry_date = formatDateString(anl_coal_rank.Expiry_date)
	anl_coal_rank.Row_effective_date = formatDateString(anl_coal_rank.Row_effective_date)
	anl_coal_rank.Row_expiry_date = formatDateString(anl_coal_rank.Row_expiry_date)






	rows, err := stmt.Exec(anl_coal_rank.Coal_rank_id, anl_coal_rank.Abbreviation, anl_coal_rank.Active_ind, anl_coal_rank.Coal_rank_scheme_id, anl_coal_rank.Effective_date, anl_coal_rank.Expiry_date, anl_coal_rank.Long_name, anl_coal_rank.Max_value, anl_coal_rank.Max_value_ouom, anl_coal_rank.Max_value_uom, anl_coal_rank.Min_value, anl_coal_rank.Min_value_ouom, anl_coal_rank.Min_value_uom, anl_coal_rank.Ppdm_guid, anl_coal_rank.Remark, anl_coal_rank.Short_name, anl_coal_rank.Source, anl_coal_rank.Row_changed_by, anl_coal_rank.Row_changed_date, anl_coal_rank.Row_created_by, anl_coal_rank.Row_created_date, anl_coal_rank.Row_effective_date, anl_coal_rank.Row_expiry_date, anl_coal_rank.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlCoalRank(c *fiber.Ctx) error {
	var anl_coal_rank dto.Anl_coal_rank

	setDefaults(&anl_coal_rank)

	if err := json.Unmarshal(c.Body(), &anl_coal_rank); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_coal_rank.Coal_rank_id = id

    if anl_coal_rank.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_coal_rank set abbreviation = :1, active_ind = :2, coal_rank_scheme_id = :3, effective_date = :4, expiry_date = :5, long_name = :6, max_value = :7, max_value_ouom = :8, max_value_uom = :9, min_value = :10, min_value_ouom = :11, min_value_uom = :12, ppdm_guid = :13, remark = :14, short_name = :15, source = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where coal_rank_id = :24")
	if err != nil {
		return err
	}

	anl_coal_rank.Row_changed_date = formatDate(anl_coal_rank.Row_changed_date)
	anl_coal_rank.Effective_date = formatDateString(anl_coal_rank.Effective_date)
	anl_coal_rank.Expiry_date = formatDateString(anl_coal_rank.Expiry_date)
	anl_coal_rank.Row_effective_date = formatDateString(anl_coal_rank.Row_effective_date)
	anl_coal_rank.Row_expiry_date = formatDateString(anl_coal_rank.Row_expiry_date)






	rows, err := stmt.Exec(anl_coal_rank.Abbreviation, anl_coal_rank.Active_ind, anl_coal_rank.Coal_rank_scheme_id, anl_coal_rank.Effective_date, anl_coal_rank.Expiry_date, anl_coal_rank.Long_name, anl_coal_rank.Max_value, anl_coal_rank.Max_value_ouom, anl_coal_rank.Max_value_uom, anl_coal_rank.Min_value, anl_coal_rank.Min_value_ouom, anl_coal_rank.Min_value_uom, anl_coal_rank.Ppdm_guid, anl_coal_rank.Remark, anl_coal_rank.Short_name, anl_coal_rank.Source, anl_coal_rank.Row_changed_by, anl_coal_rank.Row_changed_date, anl_coal_rank.Row_created_by, anl_coal_rank.Row_effective_date, anl_coal_rank.Row_expiry_date, anl_coal_rank.Row_quality, anl_coal_rank.Coal_rank_id)
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

func PatchAnlCoalRank(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_coal_rank set "
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
	query += " where coal_rank_id = :id"

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

func DeleteAnlCoalRank(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_coal_rank dto.Anl_coal_rank
	anl_coal_rank.Coal_rank_id = id

	stmt, err := db.Prepare("delete from anl_coal_rank where coal_rank_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_coal_rank.Coal_rank_id)
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


