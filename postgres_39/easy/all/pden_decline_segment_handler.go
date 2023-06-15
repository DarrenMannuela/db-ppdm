package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenDeclineSegment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_decline_segment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_decline_segment

	for rows.Next() {
		var pden_decline_segment dto.Pden_decline_segment
		if err := rows.Scan(&pden_decline_segment.Pden_subtype, &pden_decline_segment.Pden_id, &pden_decline_segment.Pden_source, &pden_decline_segment.Case_id, &pden_decline_segment.Segment_id, &pden_decline_segment.Active_ind, &pden_decline_segment.Decline_curve_type, &pden_decline_segment.Decline_type, &pden_decline_segment.Duration, &pden_decline_segment.Duration_ouom, &pden_decline_segment.Effective_date, &pden_decline_segment.End_date, &pden_decline_segment.Expiry_date, &pden_decline_segment.Final_decline, &pden_decline_segment.Final_rate, &pden_decline_segment.Initial_decline, &pden_decline_segment.Initial_rate, &pden_decline_segment.Minimum_decline, &pden_decline_segment.N_factor, &pden_decline_segment.Ppdm_guid, &pden_decline_segment.Product_type, &pden_decline_segment.Rate_ouom, &pden_decline_segment.Ratio_curve_type, &pden_decline_segment.Ratio_final_rate, &pden_decline_segment.Ratio_fluid_type, &pden_decline_segment.Ratio_initial_rate, &pden_decline_segment.Ratio_rate_ouom, &pden_decline_segment.Ratio_volume, &pden_decline_segment.Ratio_volume_ouom, &pden_decline_segment.Remark, &pden_decline_segment.Source, &pden_decline_segment.Start_date, &pden_decline_segment.Volume, &pden_decline_segment.Volume_ouom, &pden_decline_segment.Row_changed_by, &pden_decline_segment.Row_changed_date, &pden_decline_segment.Row_created_by, &pden_decline_segment.Row_created_date, &pden_decline_segment.Row_effective_date, &pden_decline_segment.Row_expiry_date, &pden_decline_segment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_decline_segment to result
		result = append(result, pden_decline_segment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenDeclineSegment(c *fiber.Ctx) error {
	var pden_decline_segment dto.Pden_decline_segment

	setDefaults(&pden_decline_segment)

	if err := json.Unmarshal(c.Body(), &pden_decline_segment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_decline_segment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41)")
	if err != nil {
		return err
	}
	pden_decline_segment.Row_created_date = formatDate(pden_decline_segment.Row_created_date)
	pden_decline_segment.Row_changed_date = formatDate(pden_decline_segment.Row_changed_date)
	pden_decline_segment.Effective_date = formatDateString(pden_decline_segment.Effective_date)
	pden_decline_segment.End_date = formatDateString(pden_decline_segment.End_date)
	pden_decline_segment.Expiry_date = formatDateString(pden_decline_segment.Expiry_date)
	pden_decline_segment.Start_date = formatDateString(pden_decline_segment.Start_date)
	pden_decline_segment.Row_effective_date = formatDateString(pden_decline_segment.Row_effective_date)
	pden_decline_segment.Row_expiry_date = formatDateString(pden_decline_segment.Row_expiry_date)






	rows, err := stmt.Exec(pden_decline_segment.Pden_subtype, pden_decline_segment.Pden_id, pden_decline_segment.Pden_source, pden_decline_segment.Case_id, pden_decline_segment.Segment_id, pden_decline_segment.Active_ind, pden_decline_segment.Decline_curve_type, pden_decline_segment.Decline_type, pden_decline_segment.Duration, pden_decline_segment.Duration_ouom, pden_decline_segment.Effective_date, pden_decline_segment.End_date, pden_decline_segment.Expiry_date, pden_decline_segment.Final_decline, pden_decline_segment.Final_rate, pden_decline_segment.Initial_decline, pden_decline_segment.Initial_rate, pden_decline_segment.Minimum_decline, pden_decline_segment.N_factor, pden_decline_segment.Ppdm_guid, pden_decline_segment.Product_type, pden_decline_segment.Rate_ouom, pden_decline_segment.Ratio_curve_type, pden_decline_segment.Ratio_final_rate, pden_decline_segment.Ratio_fluid_type, pden_decline_segment.Ratio_initial_rate, pden_decline_segment.Ratio_rate_ouom, pden_decline_segment.Ratio_volume, pden_decline_segment.Ratio_volume_ouom, pden_decline_segment.Remark, pden_decline_segment.Source, pden_decline_segment.Start_date, pden_decline_segment.Volume, pden_decline_segment.Volume_ouom, pden_decline_segment.Row_changed_by, pden_decline_segment.Row_changed_date, pden_decline_segment.Row_created_by, pden_decline_segment.Row_created_date, pden_decline_segment.Row_effective_date, pden_decline_segment.Row_expiry_date, pden_decline_segment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenDeclineSegment(c *fiber.Ctx) error {
	var pden_decline_segment dto.Pden_decline_segment

	setDefaults(&pden_decline_segment)

	if err := json.Unmarshal(c.Body(), &pden_decline_segment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_decline_segment.Pden_subtype = id

    if pden_decline_segment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_decline_segment set pden_id = :1, pden_source = :2, case_id = :3, segment_id = :4, active_ind = :5, decline_curve_type = :6, decline_type = :7, duration = :8, duration_ouom = :9, effective_date = :10, end_date = :11, expiry_date = :12, final_decline = :13, final_rate = :14, initial_decline = :15, initial_rate = :16, minimum_decline = :17, n_factor = :18, ppdm_guid = :19, product_type = :20, rate_ouom = :21, ratio_curve_type = :22, ratio_final_rate = :23, ratio_fluid_type = :24, ratio_initial_rate = :25, ratio_rate_ouom = :26, ratio_volume = :27, ratio_volume_ouom = :28, remark = :29, source = :30, start_date = :31, volume = :32, volume_ouom = :33, row_changed_by = :34, row_changed_date = :35, row_created_by = :36, row_effective_date = :37, row_expiry_date = :38, row_quality = :39 where pden_subtype = :41")
	if err != nil {
		return err
	}

	pden_decline_segment.Row_changed_date = formatDate(pden_decline_segment.Row_changed_date)
	pden_decline_segment.Effective_date = formatDateString(pden_decline_segment.Effective_date)
	pden_decline_segment.End_date = formatDateString(pden_decline_segment.End_date)
	pden_decline_segment.Expiry_date = formatDateString(pden_decline_segment.Expiry_date)
	pden_decline_segment.Start_date = formatDateString(pden_decline_segment.Start_date)
	pden_decline_segment.Row_effective_date = formatDateString(pden_decline_segment.Row_effective_date)
	pden_decline_segment.Row_expiry_date = formatDateString(pden_decline_segment.Row_expiry_date)






	rows, err := stmt.Exec(pden_decline_segment.Pden_id, pden_decline_segment.Pden_source, pden_decline_segment.Case_id, pden_decline_segment.Segment_id, pden_decline_segment.Active_ind, pden_decline_segment.Decline_curve_type, pden_decline_segment.Decline_type, pden_decline_segment.Duration, pden_decline_segment.Duration_ouom, pden_decline_segment.Effective_date, pden_decline_segment.End_date, pden_decline_segment.Expiry_date, pden_decline_segment.Final_decline, pden_decline_segment.Final_rate, pden_decline_segment.Initial_decline, pden_decline_segment.Initial_rate, pden_decline_segment.Minimum_decline, pden_decline_segment.N_factor, pden_decline_segment.Ppdm_guid, pden_decline_segment.Product_type, pden_decline_segment.Rate_ouom, pden_decline_segment.Ratio_curve_type, pden_decline_segment.Ratio_final_rate, pden_decline_segment.Ratio_fluid_type, pden_decline_segment.Ratio_initial_rate, pden_decline_segment.Ratio_rate_ouom, pden_decline_segment.Ratio_volume, pden_decline_segment.Ratio_volume_ouom, pden_decline_segment.Remark, pden_decline_segment.Source, pden_decline_segment.Start_date, pden_decline_segment.Volume, pden_decline_segment.Volume_ouom, pden_decline_segment.Row_changed_by, pden_decline_segment.Row_changed_date, pden_decline_segment.Row_created_by, pden_decline_segment.Row_effective_date, pden_decline_segment.Row_expiry_date, pden_decline_segment.Row_quality, pden_decline_segment.Pden_subtype)
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

func PatchPdenDeclineSegment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_decline_segment set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where pden_subtype = :id"

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

func DeletePdenDeclineSegment(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_decline_segment dto.Pden_decline_segment
	pden_decline_segment.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_decline_segment where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_decline_segment.Pden_subtype)
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


