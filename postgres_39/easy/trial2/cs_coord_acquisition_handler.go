package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsCoordAcquisition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_coord_acquisition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_coord_acquisition

	for rows.Next() {
		var cs_coord_acquisition dto.Cs_coord_acquisition
		if err := rows.Scan(&cs_coord_acquisition.Acquisition_id, &cs_coord_acquisition.Acquired_by_ba_id, &cs_coord_acquisition.Active_ind, &cs_coord_acquisition.Capture_method, &cs_coord_acquisition.Compute_method, &cs_coord_acquisition.Coordinate_quality, &cs_coord_acquisition.Data_create_date, &cs_coord_acquisition.Digitized_scale, &cs_coord_acquisition.Effective_date, &cs_coord_acquisition.Expiry_date, &cs_coord_acquisition.Horizontal_accuracy, &cs_coord_acquisition.Horizontal_accuracy_ouom, &cs_coord_acquisition.Ppdm_guid, &cs_coord_acquisition.Remark, &cs_coord_acquisition.Source, &cs_coord_acquisition.Source_scale, &cs_coord_acquisition.Survey_accuracy, &cs_coord_acquisition.Transform_id, &cs_coord_acquisition.Vertical_accuracy, &cs_coord_acquisition.Vertical_accuracy_ouom, &cs_coord_acquisition.Row_changed_by, &cs_coord_acquisition.Row_changed_date, &cs_coord_acquisition.Row_created_by, &cs_coord_acquisition.Row_created_date, &cs_coord_acquisition.Row_effective_date, &cs_coord_acquisition.Row_expiry_date, &cs_coord_acquisition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_coord_acquisition to result
		result = append(result, cs_coord_acquisition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsCoordAcquisition(c *fiber.Ctx) error {
	var cs_coord_acquisition dto.Cs_coord_acquisition

	setDefaults(&cs_coord_acquisition)

	if err := json.Unmarshal(c.Body(), &cs_coord_acquisition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_coord_acquisition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	cs_coord_acquisition.Row_created_date = formatDate(cs_coord_acquisition.Row_created_date)
	cs_coord_acquisition.Row_changed_date = formatDate(cs_coord_acquisition.Row_changed_date)
	cs_coord_acquisition.Data_create_date = formatDateString(cs_coord_acquisition.Data_create_date)
	cs_coord_acquisition.Effective_date = formatDateString(cs_coord_acquisition.Effective_date)
	cs_coord_acquisition.Expiry_date = formatDateString(cs_coord_acquisition.Expiry_date)
	cs_coord_acquisition.Row_effective_date = formatDateString(cs_coord_acquisition.Row_effective_date)
	cs_coord_acquisition.Row_expiry_date = formatDateString(cs_coord_acquisition.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_acquisition.Acquisition_id, cs_coord_acquisition.Acquired_by_ba_id, cs_coord_acquisition.Active_ind, cs_coord_acquisition.Capture_method, cs_coord_acquisition.Compute_method, cs_coord_acquisition.Coordinate_quality, cs_coord_acquisition.Data_create_date, cs_coord_acquisition.Digitized_scale, cs_coord_acquisition.Effective_date, cs_coord_acquisition.Expiry_date, cs_coord_acquisition.Horizontal_accuracy, cs_coord_acquisition.Horizontal_accuracy_ouom, cs_coord_acquisition.Ppdm_guid, cs_coord_acquisition.Remark, cs_coord_acquisition.Source, cs_coord_acquisition.Source_scale, cs_coord_acquisition.Survey_accuracy, cs_coord_acquisition.Transform_id, cs_coord_acquisition.Vertical_accuracy, cs_coord_acquisition.Vertical_accuracy_ouom, cs_coord_acquisition.Row_changed_by, cs_coord_acquisition.Row_changed_date, cs_coord_acquisition.Row_created_by, cs_coord_acquisition.Row_created_date, cs_coord_acquisition.Row_effective_date, cs_coord_acquisition.Row_expiry_date, cs_coord_acquisition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsCoordAcquisition(c *fiber.Ctx) error {
	var cs_coord_acquisition dto.Cs_coord_acquisition

	setDefaults(&cs_coord_acquisition)

	if err := json.Unmarshal(c.Body(), &cs_coord_acquisition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_coord_acquisition.Acquisition_id = id

    if cs_coord_acquisition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_coord_acquisition set acquired_by_ba_id = :1, active_ind = :2, capture_method = :3, compute_method = :4, coordinate_quality = :5, data_create_date = :6, digitized_scale = :7, effective_date = :8, expiry_date = :9, horizontal_accuracy = :10, horizontal_accuracy_ouom = :11, ppdm_guid = :12, remark = :13, source = :14, source_scale = :15, survey_accuracy = :16, transform_id = :17, vertical_accuracy = :18, vertical_accuracy_ouom = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where acquisition_id = :27")
	if err != nil {
		return err
	}

	cs_coord_acquisition.Row_changed_date = formatDate(cs_coord_acquisition.Row_changed_date)
	cs_coord_acquisition.Data_create_date = formatDateString(cs_coord_acquisition.Data_create_date)
	cs_coord_acquisition.Effective_date = formatDateString(cs_coord_acquisition.Effective_date)
	cs_coord_acquisition.Expiry_date = formatDateString(cs_coord_acquisition.Expiry_date)
	cs_coord_acquisition.Row_effective_date = formatDateString(cs_coord_acquisition.Row_effective_date)
	cs_coord_acquisition.Row_expiry_date = formatDateString(cs_coord_acquisition.Row_expiry_date)






	rows, err := stmt.Exec(cs_coord_acquisition.Acquired_by_ba_id, cs_coord_acquisition.Active_ind, cs_coord_acquisition.Capture_method, cs_coord_acquisition.Compute_method, cs_coord_acquisition.Coordinate_quality, cs_coord_acquisition.Data_create_date, cs_coord_acquisition.Digitized_scale, cs_coord_acquisition.Effective_date, cs_coord_acquisition.Expiry_date, cs_coord_acquisition.Horizontal_accuracy, cs_coord_acquisition.Horizontal_accuracy_ouom, cs_coord_acquisition.Ppdm_guid, cs_coord_acquisition.Remark, cs_coord_acquisition.Source, cs_coord_acquisition.Source_scale, cs_coord_acquisition.Survey_accuracy, cs_coord_acquisition.Transform_id, cs_coord_acquisition.Vertical_accuracy, cs_coord_acquisition.Vertical_accuracy_ouom, cs_coord_acquisition.Row_changed_by, cs_coord_acquisition.Row_changed_date, cs_coord_acquisition.Row_created_by, cs_coord_acquisition.Row_effective_date, cs_coord_acquisition.Row_expiry_date, cs_coord_acquisition.Row_quality, cs_coord_acquisition.Acquisition_id)
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

func PatchCsCoordAcquisition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_coord_acquisition set "
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
		} else if key == "data_create_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where acquisition_id = :id"

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

func DeleteCsCoordAcquisition(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_coord_acquisition dto.Cs_coord_acquisition
	cs_coord_acquisition.Acquisition_id = id

	stmt, err := db.Prepare("delete from cs_coord_acquisition where acquisition_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_coord_acquisition.Acquisition_id)
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


