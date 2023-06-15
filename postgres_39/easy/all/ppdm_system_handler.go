package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSystem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_system")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_system

	for rows.Next() {
		var ppdm_system dto.Ppdm_system
		if err := rows.Scan(&ppdm_system.System_id, &ppdm_system.Active_ind, &ppdm_system.Business_owner_ba_id, &ppdm_system.Connect_string, &ppdm_system.Creator_ba_id, &ppdm_system.Effective_date, &ppdm_system.Expiry_date, &ppdm_system.Operating_system, &ppdm_system.Ppdm_guid, &ppdm_system.Ppdm_system_type, &ppdm_system.Rdbms_id, &ppdm_system.Remark, &ppdm_system.Source, &ppdm_system.System_long_name, &ppdm_system.Technical_owner_ba_id, &ppdm_system.Version_num, &ppdm_system.Row_changed_by, &ppdm_system.Row_changed_date, &ppdm_system.Row_created_by, &ppdm_system.Row_created_date, &ppdm_system.Row_effective_date, &ppdm_system.Row_expiry_date, &ppdm_system.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_system to result
		result = append(result, ppdm_system)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSystem(c *fiber.Ctx) error {
	var ppdm_system dto.Ppdm_system

	setDefaults(&ppdm_system)

	if err := json.Unmarshal(c.Body(), &ppdm_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_system values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	ppdm_system.Row_created_date = formatDate(ppdm_system.Row_created_date)
	ppdm_system.Row_changed_date = formatDate(ppdm_system.Row_changed_date)
	ppdm_system.Effective_date = formatDateString(ppdm_system.Effective_date)
	ppdm_system.Expiry_date = formatDateString(ppdm_system.Expiry_date)
	ppdm_system.Row_effective_date = formatDateString(ppdm_system.Row_effective_date)
	ppdm_system.Row_expiry_date = formatDateString(ppdm_system.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system.System_id, ppdm_system.Active_ind, ppdm_system.Business_owner_ba_id, ppdm_system.Connect_string, ppdm_system.Creator_ba_id, ppdm_system.Effective_date, ppdm_system.Expiry_date, ppdm_system.Operating_system, ppdm_system.Ppdm_guid, ppdm_system.Ppdm_system_type, ppdm_system.Rdbms_id, ppdm_system.Remark, ppdm_system.Source, ppdm_system.System_long_name, ppdm_system.Technical_owner_ba_id, ppdm_system.Version_num, ppdm_system.Row_changed_by, ppdm_system.Row_changed_date, ppdm_system.Row_created_by, ppdm_system.Row_created_date, ppdm_system.Row_effective_date, ppdm_system.Row_expiry_date, ppdm_system.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSystem(c *fiber.Ctx) error {
	var ppdm_system dto.Ppdm_system

	setDefaults(&ppdm_system)

	if err := json.Unmarshal(c.Body(), &ppdm_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_system.System_id = id

    if ppdm_system.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_system set active_ind = :1, business_owner_ba_id = :2, connect_string = :3, creator_ba_id = :4, effective_date = :5, expiry_date = :6, operating_system = :7, ppdm_guid = :8, ppdm_system_type = :9, rdbms_id = :10, remark = :11, source = :12, system_long_name = :13, technical_owner_ba_id = :14, version_num = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where system_id = :23")
	if err != nil {
		return err
	}

	ppdm_system.Row_changed_date = formatDate(ppdm_system.Row_changed_date)
	ppdm_system.Effective_date = formatDateString(ppdm_system.Effective_date)
	ppdm_system.Expiry_date = formatDateString(ppdm_system.Expiry_date)
	ppdm_system.Row_effective_date = formatDateString(ppdm_system.Row_effective_date)
	ppdm_system.Row_expiry_date = formatDateString(ppdm_system.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system.Active_ind, ppdm_system.Business_owner_ba_id, ppdm_system.Connect_string, ppdm_system.Creator_ba_id, ppdm_system.Effective_date, ppdm_system.Expiry_date, ppdm_system.Operating_system, ppdm_system.Ppdm_guid, ppdm_system.Ppdm_system_type, ppdm_system.Rdbms_id, ppdm_system.Remark, ppdm_system.Source, ppdm_system.System_long_name, ppdm_system.Technical_owner_ba_id, ppdm_system.Version_num, ppdm_system.Row_changed_by, ppdm_system.Row_changed_date, ppdm_system.Row_created_by, ppdm_system.Row_effective_date, ppdm_system.Row_expiry_date, ppdm_system.Row_quality, ppdm_system.System_id)
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

func PatchPpdmSystem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_system set "
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
	query += " where system_id = :id"

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

func DeletePpdmSystem(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_system dto.Ppdm_system
	ppdm_system.System_id = id

	stmt, err := db.Prepare("delete from ppdm_system where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_system.System_id)
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


