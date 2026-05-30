class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> values = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int num = nums[i];
            int diff = target - num;
            if (values.containsKey(diff))
                return new int[] {values.get(diff), i};
            values.put(num, i);
        }
        
        return new int[] {};
    }
}
