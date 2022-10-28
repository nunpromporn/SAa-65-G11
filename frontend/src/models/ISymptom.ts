import {PatientInterface} from "./IPatient";
import {UserInterface} from "./IUser";
import {TendernessInterface} from "./ITenderness";
import {DepartmentInterface} from "./IDepartment";

export interface SymptomInterface{
    ID: number,
    Explain:   string,
    SymptomTime:    Date,

    UserID:    number,
    User:      UserInterface,

    PatientID:    number,
    Patient:      PatientInterface,

    TendernessID:   number,
    Tenderness:     TendernessInterface,

    DepartmentID:   number,
    Department:     DepartmentInterface,

}
