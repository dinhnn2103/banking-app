import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-account-balance',
  templateUrl: './account-balance.component.html',
  styleUrls: ['./account-balance.component.scss']
})
export class AccountBalanceComponent implements OnInit {

  @Input() user: any;
  date = null;
  day = null;
  month = null;
  year = null;

  constructor() { }

  ngOnInit(): void {
    this.date = new Date();
    this.day = String(this.date.getDate()).padStart(2, '0');
    this.month = String(Number(this.date.getMonth()) + 1).padStart(2, '0');
    this.year = this.date.getFullYear();
  }

}
