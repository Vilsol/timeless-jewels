export type Filter = {
  id: string;
  value: { min: number; max: number };
  disabled?: boolean;
};
export type FilterGroup = {
  type: string;
  value: { min: number };
  filters: Filter[];
  disabled: boolean;
};
export type Query = {
  query: {
    status: {
      option: string;
    };
    stats: FilterGroup[];
  };
  sort: {
    price: string;
  };
};

export const filtersToFilterGroup = (filters: Filter[], disabled: boolean): FilterGroup => ({
  type: 'count',
  value: { min: 1 },
  filters: filters,
  disabled: disabled
});

export const filterGroupsToQuery = (FilterGroups: FilterGroup[]): Query => ({
  query: {
    status: {
      option: 'online'
    },
    stats: FilterGroups
  },
  sort: {
    price: 'asc'
  }
});

export const openQueryTrade = (query: Query) => {
  const url = new URL('https://www.pathofexile.com/trade/search/Necropolis');
  url.searchParams.set('q', JSON.stringify(query));
  window.open(url, '_blank');
};
