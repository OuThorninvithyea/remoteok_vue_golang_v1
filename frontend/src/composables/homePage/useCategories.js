import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useCategories() {
    const categories = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/categories')
            categories.value = res.data
        } catch (e) {
            error.value = 'Failed to load categories'
        } finally {
            loading.value = false
        }
    })
    return { categories, loading, error }
}
