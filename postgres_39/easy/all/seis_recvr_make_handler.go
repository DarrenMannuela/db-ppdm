package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisRecvrMake(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_recvr_make")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_recvr_make

	for rows.Next() {
		var seis_recvr_make dto.Seis_recvr_make
		if err := rows.Scan(&seis_recvr_make.Seis_rcvr_make, &seis_recvr_make.Abbreviation, &seis_recvr_make.Active_ind, &seis_recvr_make.Base_horizontal_freq, &seis_recvr_make.Base_vertical_freq, &seis_recvr_make.Effective_date, &seis_recvr_make.Expiry_date, &seis_recvr_make.Long_name, &seis_recvr_make.Ppdm_guid, &seis_recvr_make.Remark, &seis_recvr_make.Seis_rcvr_type, &seis_recvr_make.Short_name, &seis_recvr_make.Source, &seis_recvr_make.Row_changed_by, &seis_recvr_make.Row_changed_date, &seis_recvr_make.Row_created_by, &seis_recvr_make.Row_created_date, &seis_recvr_make.Row_effective_date, &seis_recvr_make.Row_expiry_date, &seis_recvr_make.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_recvr_make to result
		result = append(result, seis_recvr_make)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisRecvrMake(c *fiber.Ctx) error {
	var seis_recvr_make dto.Seis_recvr_make

	setDefaults(&seis_recvr_make)

	if err := json.Unmarshal(c.Body(), &seis_recvr_make); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_recvr_make values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	seis_recvr_make.Row_created_date = formatDate(seis_recvr_make.Row_created_date)
	seis_recvr_make.Row_changed_date = formatDate(seis_recvr_make.Row_changed_date)
	seis_recvr_make.Effective_date = formatDateString(seis_recvr_make.Effective_date)
	seis_recvr_make.Expiry_date = formatDateString(seis_recvr_make.Expiry_date)
	seis_recvr_make.Row_effective_date = formatDateString(seis_recvr_make.Row_effective_date)
	seis_recvr_make.Row_expiry_date = formatDateString(seis_recvr_make.Row_expiry_date)






	rows, err := stmt.Exec(seis_recvr_make.Seis_rcvr_make, seis_recvr_make.Abbreviation, seis_recvr_make.Active_ind, seis_recvr_make.Base_horizontal_freq, seis_recvr_make.Base_vertical_freq, seis_recvr_make.Effective_date, seis_recvr_make.Expiry_date, seis_recvr_make.Long_name, seis_recvr_make.Ppdm_guid, seis_recvr_make.Remark, seis_recvr_make.Seis_rcvr_type, seis_recvr_make.Short_name, seis_recvr_make.Source, seis_recvr_make.Row_changed_by, seis_recvr_make.Row_changed_date, seis_recvr_make.Row_created_by, seis_recvr_make.Row_created_date, seis_recvr_make.Row_effective_date, seis_recvr_make.Row_expiry_date, seis_recvr_make.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisRecvrMake(c *fiber.Ctx) error {
	var seis_recvr_make dto.Seis_recvr_make

	setDefaults(&seis_recvr_make)

	if err := json.Unmarshal(c.Body(), &seis_recvr_make); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_recvr_make.Seis_rcvr_make = id

    if seis_recvr_make.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_recvr_make set abbreviation = :1, active_ind = :2, base_horizontal_freq = :3, base_vertical_freq = :4, effective_date = :5, expiry_date = :6, long_name = :7, ppdm_guid = :8, remark = :9, seis_rcvr_type = :10, short_name = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where seis_rcvr_make = :20")
	if err != nil {
		return err
	}

	seis_recvr_make.Row_changed_date = formatDate(seis_recvr_make.Row_changed_date)
	seis_recvr_make.Effective_date = formatDateString(seis_recvr_make.Effective_date)
	seis_recvr_make.Expiry_date = formatDateString(seis_recvr_make.Expiry_date)
	seis_recvr_make.Row_effective_date = formatDateString(seis_recvr_make.Row_effective_date)
	seis_recvr_make.Row_expiry_date = formatDateString(seis_recvr_make.Row_expiry_date)






	rows, err := stmt.Exec(seis_recvr_make.Abbreviation, seis_recvr_make.Active_ind, seis_recvr_make.Base_horizontal_freq, seis_recvr_make.Base_vertical_freq, seis_recvr_make.Effective_date, seis_recvr_make.Expiry_date, seis_recvr_make.Long_name, seis_recvr_make.Ppdm_guid, seis_recvr_make.Remark, seis_recvr_make.Seis_rcvr_type, seis_recvr_make.Short_name, seis_recvr_make.Source, seis_recvr_make.Row_changed_by, seis_recvr_make.Row_changed_date, seis_recvr_make.Row_created_by, seis_recvr_make.Row_effective_date, seis_recvr_make.Row_expiry_date, seis_recvr_make.Row_quality, seis_recvr_make.Seis_rcvr_make)
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

func PatchSeisRecvrMake(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_recvr_make set "
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
	query += " where seis_rcvr_make = :id"

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

func DeleteSeisRecvrMake(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_recvr_make dto.Seis_recvr_make
	seis_recvr_make.Seis_rcvr_make = id

	stmt, err := db.Prepare("delete from seis_recvr_make where seis_rcvr_make = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_recvr_make.Seis_rcvr_make)
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


